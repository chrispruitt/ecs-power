resource "aws_lambda_function" "ecs_power" {
  function_name    = "ecs-power"
  s3_bucket        = var.s3_bucket
  s3_key           = var.s3_key
  role             = aws_iam_role.lambda_role.arn
  handler          = "ecs-power"
  runtime          = "go1.x"

  lifecycle {
    ignore_changes = [
      last_modified
      ]
  }
}

variable "s3_bucket" {
  type = string
}

variable "s3_key" {
  type = string
}
