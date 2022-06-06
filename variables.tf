variable "resource_group" {
  description = "resource group options"
  type = object({
    location = string
    name     = string
  })
  validation {
    condition     = length(regexall("\\b \\b", var.resource_group.location)) == 0
    error_message = "Spaces between the words are not allowed."
  }
}

variable "user_assigned_identity_name" {
  description = "name of user identity"
  type        = string
}

