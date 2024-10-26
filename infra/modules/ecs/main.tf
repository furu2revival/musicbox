resource "aws_ecs_cluster" "main" {
  name = var.env
}

data "aws_iam_policy_document" "main-1" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "main-1" {
  name               = "ecs-iam-role-${var.env}"
  assume_role_policy = data.aws_iam_policy_document.main-1.json
}

resource "aws_iam_role_policy_attachment" "main-1" {
  role       = aws_iam_role.main-1.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_ecs_task_definition" "main" {
  family                   = var.env
  cpu                      = 1024
  memory                   = 2048
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.main-1.arn
  task_role_arn            = aws_iam_role.main-1.arn
  container_definitions = jsonencode([
    {
      "name" : var.env,
      "image" : "ghcr.io/furu2revival/musicbox/api-server:latest",
      "portMappings" : [
        {
          "containerPort" : 8000,
          "hostPort" : 8000,
        }
      ],
      "environment" : [
        {
          "name" : "MUSICBOX_CONFIG_FILEPATH",
          "value" : "/musicbox.json",
        },
      ],
      "logConfiguration" : {
        "logDriver" : "awslogs",
        "options" : {
          "awslogs-region" : var.region,
          "awslogs-stream-prefix" : "backend",
          "awslogs-group" : "logs-${var.env}",
        }
      },
    }
  ])
}

data "aws_ecs_task_definition" "main" {
  task_definition = aws_ecs_task_definition.main.family
}

resource "aws_ecs_service" "main" {
  name            = aws_ecs_cluster.main.name
  cluster         = aws_ecs_cluster.main.id
  task_definition = data.aws_ecs_task_definition.main.arn
  desired_count   = 2
  launch_type     = "FARGATE"

  network_configuration {
    subnets         = var.private_subnet_ids
    security_groups = [aws_security_group.main.id]
  }

  load_balancer {
    target_group_arn = var.lb_target_group_arn
    container_name   = var.env
    container_port   = 8000
  }
}

resource "aws_appautoscaling_target" "main" {
  service_namespace  = "ecs"
  resource_id        = "service/${aws_ecs_cluster.main.name}/${aws_ecs_service.main.name}"
  scalable_dimension = "ecs:service:DesiredCount"
  min_capacity       = 2
  max_capacity       = 4
}

resource "aws_appautoscaling_policy" "main-1" {
  name               = "ecs-scale-up-${var.env}"
  service_namespace  = aws_appautoscaling_target.main.service_namespace
  resource_id        = aws_appautoscaling_target.main.resource_id
  scalable_dimension = aws_appautoscaling_target.main.scalable_dimension

  step_scaling_policy_configuration {
    adjustment_type         = "ChangeInCapacity"
    cooldown                = 60
    metric_aggregation_type = "Average"

    step_adjustment {
      metric_interval_lower_bound = 0
      scaling_adjustment          = 1
    }
  }
}

resource "aws_appautoscaling_policy" "main-2" {
  name               = "ecs-scale-down-${var.env}"
  service_namespace  = aws_appautoscaling_target.main.service_namespace
  resource_id        = aws_appautoscaling_target.main.resource_id
  scalable_dimension = aws_appautoscaling_target.main.scalable_dimension

  step_scaling_policy_configuration {
    adjustment_type         = "ChangeInCapacity"
    cooldown                = 60
    metric_aggregation_type = "Average"

    step_adjustment {
      metric_interval_upper_bound = 0
      scaling_adjustment          = -1
    }
  }
}

resource "aws_cloudwatch_metric_alarm" "main-1" {
  alarm_name          = "ecs-high-cpu-utilization-${var.env}"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = "1"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/ECS"
  period              = "60"
  statistic           = "Average"
  threshold           = "80"
  dimensions = {
    ClusterName = aws_ecs_cluster.main.name
    ServiceName = aws_ecs_service.main.name
  }
  alarm_actions = [
    aws_appautoscaling_policy.main-1.arn
  ]
}

resource "aws_cloudwatch_metric_alarm" "main-2" {
  alarm_name          = "ecs-low-cpu-utilization-${var.env}"
  comparison_operator = "LessThanOrEqualToThreshold"
  evaluation_periods  = "1"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/ECS"
  period              = "60"
  statistic           = "Average"
  threshold           = "30"
  dimensions = {
    ClusterName = aws_ecs_cluster.main.name
    ServiceName = aws_ecs_service.main.name
  }
  alarm_actions = [
    aws_appautoscaling_policy.main-2.arn
  ]
}

resource "aws_security_group" "main" {
  name   = "ecs-sg-${var.env}"
  vpc_id = var.vpc_id
}

resource "aws_security_group_rule" "main-ingress" {
  security_group_id = aws_security_group.main.id
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 8000
  to_port           = 8000
  cidr_blocks       = ["10.0.0.0/16"]
}

resource "aws_security_group_rule" "main-egress" {
  security_group_id = aws_security_group.main.id
  type              = "egress"
  protocol          = "all"
  from_port         = 0
  to_port           = 0
  cidr_blocks       = ["0.0.0.0/0"]
}
