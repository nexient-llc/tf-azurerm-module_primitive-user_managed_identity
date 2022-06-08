resource "azurerm_user_assigned_identity" "main" {
  resource_group_name = var.resource_group.name
  location            = var.resource_group.location

  name = var.user_assigned_identity_name
}
