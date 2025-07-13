variable "name" {
  type        = string
  description = "The name of the OIDC provider"
  default     = "github-oidc-provider"
}

variable "tags" {
  type        = map(string)
  description = "The tags of the OIDC provider"
  validation {
    condition     = !contains(keys(var.tags), "ManagedBy") && !contains(keys(var.tags), "Name")
    error_message = "The tags must not contain ManagedBy or Name keys as they will be automatically added."
  }
}

variable "client_id_list" {
  type        = list(string)
  description = "The list of client IDs for the OIDC provider"
  default     = ["sts.amazonaws.com"]
}

variable "thumbprint_list" {
  type        = list(string)
  description = "The list of thumbprints for the OIDC provider"
  default = [
    "6938fd4d98bab03faadb97b34396831e3780aea1",
    "1c58a3a8518e8759bf075b76b750d4f2df264fcd"
  ]
}

variable "url" {
  type        = string
  description = "The URL of the OIDC provider"
  default     = "https://token.actions.githubusercontent.com"
}

