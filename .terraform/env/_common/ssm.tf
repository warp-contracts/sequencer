resource "aws_ssm_parameter" "ecr_name" {
  name  = "/${var.env}/redstone/warp-contract/${local.service-name}/docker-repository"
  type  = "String"
  value = aws_ecrpublic_repository.this.repository_uri
}
