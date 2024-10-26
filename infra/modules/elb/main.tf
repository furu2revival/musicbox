resource "aws_lb" "main" {
  name                       = var.env
  load_balancer_type         = "application"
  security_groups = [aws_security_group.main.id]
  subnets                    = var.public_subnet_ids
  enable_deletion_protection = true
}

resource "aws_lb_listener" "main" {
  load_balancer_arn = aws_lb.main.arn
  protocol          = "HTTP"

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      message_body = "Not found"
      status_code  = "404"
    }
  }
}

resource "aws_lb_listener_rule" "main" {
  listener_arn = aws_lb_listener.main.arn

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.main.id
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_lb_target_group" "main" {
  name        = "lb-target-group-${var.env}"
  port        = "80"
  protocol    = "HTTP"
  vpc_id      = var.vpc_id
  target_type = "ip"

  health_check {
    port = "8000"
    path = "/grpc.health.v1.Health/Check"
  }
}

resource "aws_security_group" "main" {
  name   = "lb-sg-${var.env}"
  vpc_id = var.vpc_id
}

resource "aws_security_group_rule" "main-ingress" {
  security_group_id = aws_security_group.main.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 80
  to_port           = 80
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "main-egress" {
  security_group_id = aws_security_group.main.id
  type              = "egress"
  protocol          = "all"
  from_port         = 0
  to_port           = 0
  cidr_blocks = ["0.0.0.0/0"]
}
