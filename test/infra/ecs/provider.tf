provider "aws" {
    region = "ap-northeast-2"
    profile = "default"
}

# Uncomment this to use S3 backend. 
# If you want to use with multiple users, then use dynamodb lock.

# terraform {
#     backend "s3" {
#         bucket = "terraform-state-bucket"
#         key = "terraform.tfstate"
#         region = "ap-northeast-2"
#     }
# }

# Uncomment this to use dynamodb lock.
# Dynamodb lock is recommended to use with multiple users.
# Dynamodb should be created manually or via cloudformation before using this.

# DynamoDB Table Structure
# mandatory:
#   Partition key: LockID (string)
# recommended:
#   Capacity mode: On-demand
#   Billing mode: Pay-per-request

# terraform {
#     backend "s3" {
#         bucket = "terraform-state-bucket"
#         key = "terraform.tfstate"
#         region = "ap-northeast-2"
#         dynamodb_table = "terraform-lock"
#         encrypt = true
#     }
# }
