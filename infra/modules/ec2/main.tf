resource "aws_eip" "main" {
  domain = "vpc"
}

resource "aws_eip_association" "main" {
  instance_id   = aws_instance.main.id
  allocation_id = aws_eip.main.id
}

resource "aws_instance" "main" {
  ami                         = "ami-0ffac3e16de16665e"
  instance_type               = var.instance_type
  subnet_id                   = var.public_subnet_id
  associate_public_ip_address = true
  vpc_security_group_ids      = [aws_security_group.main.id]

  lifecycle {
    ignore_changes = [user_data]
  }

  tags = {
    Name = "bastion-${var.env}"
  }
}

resource "aws_security_group" "main" {
  name   = "bastion-${var.env}"
  vpc_id = var.vpc_id
}

resource "aws_security_group_rule" "main-ingress" {
  security_group_id = aws_security_group.main.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 22
  to_port           = 22
  // まぁ・・このままでいいか
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "main-egress" {
  security_group_id = aws_security_group.main.id
  type              = "egress"
  protocol          = "all"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
}
