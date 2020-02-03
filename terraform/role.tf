data "aws_iam_policy_document" "lambda_policy_doc" {
  statement {
    effect    = "Allow"
    actions   = [
      "logs:CreateLogStream",
      "logs:DescribeLogGroups",
      "logs:DescribeLogStreams",
      "logs:CreateLogGroup",
      "logs:PutLogEvents",
      "ssm:Get*",
      "ec2:RunInstances",
      "autoscaling:UpdateAutoScalingGroup"
    ]
    resources = ["*"]
  }
}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda_role" {
  name               = "ecs-power"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

resource "aws_iam_policy" "lambda_policy" {
  name   = "ecs-power"
  path   = "/"
  policy = data.aws_iam_policy_document.lambda_policy_doc.json
}

//# Attach IAM policies
resource "aws_iam_role_policy_attachment" "policy_attachment" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.lambda_policy.arn
}
