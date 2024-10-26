resource "random_password" "main" {
  length           = 16
  special          = true
  override_special = "!#$%^&*()-_=+[]{}<>:?"
}

resource "aws_db_instance" "rds" {
  allocated_storage          = 20
  storage_type               = "gp2"
  engine                     = "postgres"
  engine_version             = "15.8"
  instance_class             = var.instance_type
  identifier                 = var.env
  username                   = "master"
  password                   = random_password.main.result
  skip_final_snapshot        = true
  vpc_security_group_ids     = [aws_security_group.main.id]
  db_subnet_group_name       = aws_db_subnet_group.main.name
  backup_retention_period    = 30
  deletion_protection        = false
  apply_immediately          = true
  auto_minor_version_upgrade = false

  lifecycle {
    ignore_changes = [
      password,
      availability_zone
    ]
  }
}

resource "aws_db_subnet_group" "main" {
  name       = var.env
  subnet_ids = var.private_subnet_ids
}

resource "aws_security_group" "main" {
  name   = var.env
  vpc_id = var.vpc_id
}

resource "aws_security_group_rule" "main-ingress-1" {
  security_group_id        = aws_security_group.main.id
  type                     = "ingress"
  protocol                 = "tcp"
  from_port                = "5432"
  to_port                  = "5432"
  source_security_group_id = var.bastion_server_security_group_id
}

resource "aws_security_group_rule" "main-ingress-2" {
  security_group_id = aws_security_group.main.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = "5432"
  to_port           = "5432"
  cidr_blocks       = ["10.0.0.0/16"]
}
