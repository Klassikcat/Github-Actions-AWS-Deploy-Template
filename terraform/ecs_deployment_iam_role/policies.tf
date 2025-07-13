# ECS Task Execution Role Policy
resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_policy" {
  role       = aws_iam_role.github_actions_deployment.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# CloudWatch Logs Policy
resource "aws_iam_role_policy_attachment" "cloudwatch_logs_policy" {
  role       = aws_iam_role.github_actions_deployment.name
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
}

