resource "aws_secretsmanager_secret" "main" {
  name = var.env
}

resource "aws_secretsmanager_secret_version" "main" {
  secret_id = aws_secretsmanager_secret.main.id
  secret_string = jsonencode({
    CONFIG_JSON = var.config_json
  })
}
