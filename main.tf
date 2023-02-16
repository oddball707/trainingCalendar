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


data "archive_file" "zip" {
  type        = "zip"
  source_file = "lambda/main"
  output_path = "main.zip"
}

resource "random_pet" "lambda_bucket_name" {
  prefix = "training-calendar"
  length = 4
}

resource "aws_s3_bucket" "lambda_bucket" {
  bucket = random_pet.lambda_bucket_name.id
}

resource "aws_s3_object" "lambda_training_calendar" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "main.zip"
  source = data.archive_file.zip.output_path

  etag = filemd5(data.archive_file.zip.output_path)
}

resource "aws_lambda_function" "training-calendar-generator" {
  function_name    = "time"
  s3_bucket        = aws_s3_bucket.lambda_bucket.id
  s3_key           = aws_s3_object.lambda_training_calendar.key
  handler          = "main"
  source_code_hash = "data.archive_file.zip.output_base64sha256"
  role             = aws_iam_role.iam_for_lambda.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 10
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      }
    ]
  })
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