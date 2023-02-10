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
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region  = "us-west-2"
}

module "amplify" {
  source = "aws/amplify/site"
  name   = "training-calendar-frontend"

  build_settings = {
    backend_environment_name = "prod"
  }
}

data "archive-file" "zip" {
  type        = "zip"
  source_file = "lambda/main"
  output_path = "main.zip"
}

resource "aws_lambda_function" "training-calendar-generator" {
  function_name    = "time"
  filename         = "main.zip"
  handler          = "main"
  source_code_hash = "data.archive_file.zip.output_base64sha256"
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

module "lambda" {
  source = "aws/lambda"
  name   = "training-calendar-generator"

  function-name = "training-calendar-generator"
  
  vpc_config = {
    subnet_ids = [
      module.amplify.subnet_ids
    ]
  }

  tags = {
    Terraform = "True"
    Environment = "prod"
  }
}

output "lambda_arn" {
  value = module.lambda.arn
}

output "amplify_app_id" {
  value = module.amplify.app_id
}
