output "secrets_manager_arn" {
  value = aws_secretsmanager_secret_version.main.arn
}

output "secrets_manager_id" {
  value = aws_secretsmanager_secret.main.id
}

output "secret_arn" {
  value = aws_secretsmanager_secret.main.arn
}
