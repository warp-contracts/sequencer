data "aws_region" "current" {}
data "aws_lb" "alb" {
  name = data.aws_ssm_parameter.alb-security-group-name.value
}
data "aws_ssm_parameter" "alb-name" {
  name = "/${var.env}/redstone/infrastructure/alb-name"
}
data "aws_ssm_parameter" "alb-http-listener-arn" {
  name = "/${var.env}/redstone/infrastructure/alb-http-listener-arn"
}
data "aws_ssm_parameter" "vpc-id" {
  name = "/${var.env}/redstone/infrastructure/vpc-id"
}
data "aws_ssm_parameter" "cluster-arn" {
  name = "/${var.env}/redstone/infrastructure/ecs-cluster-arn"
}
data "aws_ssm_parameter" "cluster-name" {
  name = "/${var.env}/redstone/infrastructure/ecs-cluster-name"
}
data "aws_ssm_parameter" "alb-security-group-name" {
  name = "/${var.env}/redstone/infrastructure/alb-security-group-name"
}
data "aws_ssm_parameter" "warp-rds-endpoint" {
  name = "/${var.env}/redstone/infrastructure/warp/rds/endpoint"
}
data "aws_ssm_parameter" "warp-rds-endpoint-master-username" {
  name = "/${var.env}/redstone/infrastructure/warp/rds/master-username"
}
data "aws_ssm_parameter" "warp-rds-endpoint-master-password" {
  name = "/${var.env}/redstone/infrastructure/warp/rds/master-password"
}
data "aws_ssm_parameter" "warp-arweave-key" {
  name = "/${var.env}/warp-contract/gateway/arweave-jwk-key"
}
data "aws_ssm_parameter" "warp-vrf-private" {
  name = "/${var.env}/warp-contract/gateway/vrf/private"
}
data "aws_subnets" "private" {
  filter {
    name   = "vpc-id"
    values = [data.aws_ssm_parameter.vpc-id.value]
  }
  tags = {
    Tier = "Private"
  }
}
data "aws_security_group" "lb" {
  name = data.aws_ssm_parameter.alb-security-group-name.value
}
