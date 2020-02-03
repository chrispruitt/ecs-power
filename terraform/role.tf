data "aws_iam_policy_document" "lambda_policy_doc" {
  statement {
    effect    = "Allow"
    actions   = [
      "logs:CreateLogStream",
      "logs:DescribeLogGroups",
      "logs:DescribeLogStreams",
      "logs:CreateLogGroup",
      "logs:PutLogEvents"
    ]
    resources = [
      "*"]
  }

  statement {
    effect = "Allow"

    actions = [
      "ssm:Get*",
    ]

    resources = [
      "arn:*:ssm:${data.terraform_remote_state.network.outputs.aws_region}:*:parameter/dev/ecs_cluster/*",
      "arn:*:ssm:${data.terraform_remote_state.network.outputs.aws_region}:*:parameter/tst/ecs_cluster/*",
      "arn:*:ssm:${data.terraform_remote_state.network.outputs.aws_region}:*:parameter/stg/ecs_cluster/*",
    ]
  }

  statement {
    effect    = "Allow"
    actions   = [
      "autoscaling:UpdateAutoScalingGroup",
    ]
    resources = [
      "arn:*:autoscaling:${data.terraform_remote_state.network.outputs.aws_region}:*:autoScalingGroup:*:autoScalingGroupName/dev-ecs",
      "arn:*:autoscaling:${data.terraform_remote_state.network.outputs.aws_region}:*:autoScalingGroup:*:autoScalingGroupName/tst-ecs",
      "arn:*:autoscaling:${data.terraform_remote_state.network.outputs.aws_region}:*:autoScalingGroup:*:autoScalingGroupName/stg-ecs",
    ]
  }

  statement {
    effect    = "Allow"
    actions   = [
      "ec2:RunInstances",
    ]
    resources = [
      "*"]
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
