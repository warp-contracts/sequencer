resource "aws_ecs_service" "this" {
  name        = "${local.service-name}-${var.env}"
  cluster     = data.aws_ssm_parameter.cluster-arn.value
  launch_type = "EC2"

  task_definition = aws_ecs_task_definition.this.arn
  desired_count   = 1
#  deployment_maximum_percent = 100
#  deployment_minimum_healthy_percent = 100

  load_balancer {
    container_name   = local.service-name
    container_port   = local.port
    target_group_arn = aws_lb_target_group.this.arn
  }
  network_configuration {
    subnets = data.aws_subnets.private.ids
    security_groups = [
      data.aws_security_group.lb.id
    ]
  }
}

resource "aws_ecs_task_definition" "this" {
  family                = local.service-name
  execution_role_arn    = aws_iam_role.task-execution-role.arn
  network_mode          = "awsvpc"
  container_definitions = jsonencode([
    {
      name = local.service-name
      image = local.image
      essential               = true
      cpu                     = 512
      memory                  = 2048
      memoryReservation       = 64
      RequiresCompatibilities = ["EC2"]
      healthCheck = {
        interval = 30
        retries  = 3
        timeout  = 3
        command = [
          "CMD-SHELL",
          "curl -f http://localhost:${local.port}/gateway/arweave/info || exit 1"
        ]
      }
      portMappings = [
        {
          hostPort = local.port
          containerPort = local.port
          protocol = "tcp"
        }
      ]
      containerDefinitions = [
        {
          image = local.image
        }
      ]
      environment = [
        {
          name = "SEQUENCER_PORT"
          value = tostring(local.port)
        },
        {
          name = "LOG_LEVEL"
          value = "debug"
        },
      ]
      secrets = [
        {
          name = "POSTGRES_HOST"
          valueFrom = data.aws_ssm_parameter.warp-rds-endpoint.arn
        },
        {
          name = "POSTGRES_USER"
          valueFrom = data.aws_ssm_parameter.warp-rds-endpoint-master-username.arn
        },
        {
          name = "POSTGRES_PASSWORD"
          valueFrom = data.aws_ssm_parameter.warp-rds-endpoint-master-password.arn
        },
        {
          name = "ARWEAVE_WALLETJWK"
          valueFrom = data.aws_ssm_parameter.warp-arweave-key.arn
        },
        {
          name = "VRF_PRIVATEKEY",
          valueFrom = data.aws_ssm_parameter.warp-vrf-private.arn
        }
      ]
      logConfiguration = {
        logDriver = "awslogs",
        options = {
          "awslogs-group" = aws_cloudwatch_log_group.task-logs.name
          "awslogs-region" = data.aws_region.current.name
          "awslogs-stream-prefix" = "ecs"
        }
      }
      volumesFrom       = []
      mountPoints       = []
    }
  ])
}
resource "aws_cloudwatch_log_group" "task-logs" {
  name = "/${var.env}/ecs/redstone/${local.service-name}"

  retention_in_days = 180
}
resource "aws_cloudwatch_log_metric_filter" "error-log" {
  name           = "${local.service-name}-${var.env}-errors"
  pattern        = "{ $.level = \"error\" }"
  log_group_name = aws_cloudwatch_log_group.task-logs.name

  metric_transformation {
    name      = "${aws_cloudwatch_log_group.task-logs.name}-errors"
    namespace = "/${var.env}/redstone/${local.service-name}"
    value     = "1"
  }
}
resource "aws_iam_role_policy_attachment" "ecs-task-execution-role-default-policy-attachment" {
  role       = aws_iam_role.task-execution-role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}
resource "aws_lb_target_group" "this" {
  name_prefix = substr(replace(local.service-name, "-", ""), 0, 5)
  port        = local.port
  vpc_id      = data.aws_ssm_parameter.vpc-id.value
  target_type = "ip"
  protocol    = "HTTP"
  health_check {
    enabled           = true
    healthy_threshold = 2
    path              = "/gateway/arweave/info"
  }
  lifecycle {
    create_before_destroy = true
  }
}
resource "aws_lb_listener_rule" "this" {
  listener_arn = data.aws_ssm_parameter.alb-http-listener-arn.value
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.this.arn
  }
  condition {
    path_pattern {
      values = [
        "/gateway/sequencer/*",
        "/gateway/arweave/*"
      ]
    }
  }
}
resource "aws_iam_role" "task-execution-role" {
  name               = "${local.service-name}-task-execution-role"
  assume_role_policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "ecs-tasks.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
POLICY
  inline_policy {
    name   = "allowSecretsAccess"
    policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "ssm:GetParameters",
            "Resource": [
                "*"
            ]
        }
    ]
}
POLICY
  }
}

resource "aws_iam_role" "service-role" {
  name               = "${local.service-name}-ECSService"
  assume_role_policy = data.aws_iam_policy_document.ecs-service.json
}

data "aws_iam_policy_document" "ecs-service" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs.amazonaws.com"]
    }
  }
}
