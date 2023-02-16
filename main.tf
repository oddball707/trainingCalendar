terraform {
  cloud {
    organization = "oddball707"

    workspaces {
      name = "training-calendar"
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.54"
    }
  }

  required_version = ">= 0.14.0"
}

provider "aws" {
  region  = "us-west-2"
}

resource "null_resource" "gobuild" {
  provisioner "local-exec" {
    command = "cd lambda && CGO_ENABLED=0 go build && zip -r main.zip main"
  }
}

resource "aws_lambda_function" "training-calendar-generator" {
  function_name    = "training-calendar-generator"
  filename         = "lambda/main.zip"
  handler          = "main"
  role             = "${aws_iam_role.iam_for_lambda.arn}"
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 10
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_api_gateway_rest_api" "api" {
  name = "time_api"
}

resource "aws_api_gateway_resource" "resource" {
  path_part   = "time"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
  resource_id   = "${aws_api_gateway_resource.resource.id}"
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_amplify_app" "training-calendar-frontend" {
  name       = "Training Calendar Frontend"
  repository = "github.com/oddball707/trainingCalendar"
  iam_service_role_arn = aws_iam_role.iam_for_lambda.arn
  enable_branch_auto_build = true
  build_spec = <<-EOT
    version: 0.1
    frontend:
      phases:
        preBuild:
          commands:
            - cd app
            - npm install
        build:
          commands:
            - cd app
            - npm run build
      artifacts:
        baseDirectory: app/build
        files:
          - '**/*'
      cache:
        paths:
          - app/node_modules/**/*
  EOT
  # The default rewrites and redirects added by the Amplify Console.
  custom_rule {
    source = "/<*>"
    status = "404"
    target = "/index.html"
  }
  environment_variables = {
    ENV = "dev"
  }
}