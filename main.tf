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

resource "aws_lambda_function" "training-calendar-generator" {
  function_name    = "training-calendar-generator"
  filename         = "lambda/lambda.zip"
  handler          = "lambda"
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

# Attach role to Managed Policy
variable "iam_policy_arn" {
  description = "IAM Policy to be attached to role"
  type        = list(string)

  default = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]
}

resource "aws_iam_policy_attachment" "role_attach" {
  name       = "policy-lambda"
  roles      = [aws_iam_role.iam_for_lambda.id]
  count      = length(var.iam_policy_arn)
  policy_arn = element(var.iam_policy_arn, count.index)
}

resource "aws_api_gateway_rest_api" "api" {
  name = "training_cal_api"
}

resource "aws_api_gateway_resource" "proxy" {
  path_part   = "{proxy+}"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
}

//CORS
resource "aws_api_gateway_method" "options_method" {
    rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
    resource_id   = "${aws_api_gateway_resource.proxy.id}"
    http_method   = "OPTIONS"
    authorization = "NONE"
}

resource "aws_api_gateway_method_response" "options_200" {
    rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
    resource_id   = "${aws_api_gateway_resource.proxy.id}"
    http_method   = "${aws_api_gateway_method.options_method.http_method}"
    status_code   = "200"
    response_models = {
        "application/json" = "Empty"
    }
    response_parameters = {
        "method.response.header.Access-Control-Allow-Headers" = true,
        "method.response.header.Access-Control-Allow-Methods" = true,
        "method.response.header.Access-Control-Allow-Origin" = true
    }
    depends_on = [aws_api_gateway_method.options_method]
}

resource "aws_api_gateway_integration" "options_integration" {
    rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
    resource_id   = "${aws_api_gateway_resource.proxy.id}"
    http_method   = "${aws_api_gateway_method.options_method.http_method}"
    type          = "MOCK"
    depends_on = [aws_api_gateway_method.options_method]
}

resource "aws_api_gateway_integration_response" "options_integration_response" {
    rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
    resource_id   = "${aws_api_gateway_resource.proxy.id}"
    http_method   = "${aws_api_gateway_method.options_method.http_method}"
    status_code   = "${aws_api_gateway_method_response.options_200.status_code}"
    response_parameters = {
        "method.response.header.Access-Control-Allow-Headers" = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
        "method.response.header.Access-Control-Allow-Methods" = "'GET,OPTIONS,POST,PUT'",
        "method.response.header.Access-Control-Allow-Origin" = "'*'"
    }
    depends_on = [aws_api_gateway_method_response.options_200]
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
  resource_id   = "${aws_api_gateway_resource.proxy.id}"
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "method_response_200" {
    rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
    resource_id   = "${aws_api_gateway_resource.proxy.id}"
    http_method   = "${aws_api_gateway_method.method.http_method}"
    status_code   = "200"
    response_parameters = {
        "method.response.header.Access-Control-Allow-Origin" = true
    }
    depends_on = [aws_api_gateway_method.method]
}

resource "aws_api_gateway_method" "proxy_root" {
  rest_api_id   = aws_api_gateway_rest_api.api.id
  resource_id   = aws_api_gateway_rest_api.api.root_resource_id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_method_response" "proxy_root_response_200" {
    rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
    resource_id   = "${aws_api_gateway_resource.proxy.id}"
    http_method   = "${aws_api_gateway_method.proxy_root.http_method}"
    status_code   = "200"
    response_parameters = {
        "method.response.header.Access-Control-Allow-Origin" = true
    }
    depends_on = [aws_api_gateway_method.proxy_root]
}


resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_method.method.resource_id
  http_method             = aws_api_gateway_method.method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.training-calendar-generator.invoke_arn
}

resource "aws_api_gateway_integration" "integration_root" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_method.proxy_root.resource_id
  http_method             = aws_api_gateway_method.proxy_root.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.training-calendar-generator.invoke_arn
}

resource "aws_api_gateway_deployment" "api_deployment" {
  depends_on = [
    aws_api_gateway_integration.integration,
    aws_api_gateway_integration.integration_root,
  ]

  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name  = "api"
}

resource "aws_lambda_permission" "lambda_permission" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.training-calendar-generator.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_deployment.api_deployment.execution_arn}/*/*"
}