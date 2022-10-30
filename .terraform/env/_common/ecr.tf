resource "aws_ecrpublic_repository" "this" {
  repository_name = local.service-name
  provider = aws.virginia
}
