terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_ecs_cluster" "api_cluster" {
  name = "mock-api-cluster"
}

#task definition
resource "aws_ecs_task_definition" "api_task" {
  family                   = "mock-api-task"
  container_definitions    = <<DEFINITION
  [
    {
        "name": "mock-api",
        "image": "docker.io/alikhakpouri/mock-api:1.0",
        "essential": true,
        "portMappings": [
            {
                "containerPort": 5000,
                "hostPort": 5000
            }
        ],
        "memory": 512,
        "cpu": 256
    }
  ]
  DEFINITION
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  memory                   = 512
  cpu                      = 256
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name               = "ecsTaskExecutionRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

#add the vpc
resource "aws_default_vpc" "vpc" {
}

resource "aws_default_subnet" "default_subnet_one" {
  availability_zone = "us-east-1a"

}

resource "aws_default_subnet" "default_subnet_two" {
  availability_zone = "us-east-1b"
}

#load balancer
resource "aws_alb" "application_lobad_balancer" {
  name               = "mock-api-loadbalancer"
  load_balancer_type = "application"
  subnets = [
    "${aws_default_subnet.default_subnet_one.id}",
    "${aws_default_subnet.default_subnet_two.id}"
  ]
  security_groups = ["${aws_security_group.load_balancer_security_group.id}"]
}

#security groups for the load balancer

resource "aws_security_group" "load_balancer_security_group" {
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Allow traffic in from all sources
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

#configure load balancer with vpc
resource "aws_lb_target_group" "target_group" {
  name        = "mocket-api-target-group"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_default_vpc.vpc.id
}

resource "aws_lb_listener" "listener" {
  load_balancer_arn = aws_alb.application_lobad_balancer.arn
  port              = 80
  protocol          = "HTTP"
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.target_group.arn
  }
}

#ecs service
resource "aws_ecs_service" "app_service" {
  name            = "mock-api-service"
  cluster         = aws_ecs_cluster.api_cluster.id
  task_definition = aws_ecs_task_definition.api_task.arn
  launch_type     = "FARGATE"
  desired_count   = 3
  load_balancer {
    target_group_arn = aws_lb_target_group.target_group.arn
    container_name   = aws_ecs_task_definition.api_task.family
    container_port   = 5000
  }
  network_configuration {
    subnets = [
      aws_default_subnet.default_subnet_one.id,
      aws_default_subnet.default_subnet_two.id
    ]
    assign_public_ip = true
    security_groups  = [aws_security_group.service_security_group.id]
  }
}

#Only allow the traffic from the created load balancer
resource "aws_security_group" "service_security_group" {
  ingress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    # Only allowing traffic in from the load balancer security group
    security_groups = ["${aws_security_group.load_balancer_security_group.id}"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
