data "aws_iam_policy_document" "ecs_deployment_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "github_actions_deployment" {
  name               = var.role_name
  assume_role_policy = data.aws_iam_policy_document.ecs_deployment_assume_role_policy.json
}

