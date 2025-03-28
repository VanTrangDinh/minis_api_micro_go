package enums

type PermissionResource string

const (
	ResourceUser       PermissionResource = "user"
	ResourceRole       PermissionResource = "role"
	ResourcePermission PermissionResource = "permission"
	ResourceAuth       PermissionResource = "auth"
	ResourceToken      PermissionResource = "token"
)

type PermissionAction string

const (
	ActionCreate     PermissionAction = "create"
	ActionRead       PermissionAction = "read"
	ActionUpdate     PermissionAction = "update"
	ActionDelete     PermissionAction = "delete"
	ActionAssign     PermissionAction = "assign"
	ActionRemove     PermissionAction = "remove"
	ActionActivate   PermissionAction = "activate"
	ActionDeactivate PermissionAction = "deactivate"
)
