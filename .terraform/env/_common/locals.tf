locals {
  service-name = "warp-sequencer"
  sg-name = "ecs-service-${local.service-name}-${var.env}"
  port = 80
  image = "${aws_ecrpublic_repository.this.repository_uri}:${var.image-tag}"
}
