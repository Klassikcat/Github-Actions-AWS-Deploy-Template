resource "aws_service_discovery_http_namespace" "ecs_namespace" {
  name = "ecs-cluster.local"
  tags = {
    Name = "ecs-cluster.local"
    ManagedBy = "terraform"
  }
}

resource "aws_ecs_cluster" "ecs_cluster" {
  name = "ecs-cluster"
  service_connect_defaults {
    namespace = aws_service_discovery_http_namespace.ecs_namespace.arn
  }
  tags = {
    Name = "ecs-cluster"
    ManagedBy = "terraform"
  }
}
