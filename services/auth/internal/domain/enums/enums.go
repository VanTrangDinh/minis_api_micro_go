package enums

type PermissionResource string
type PermissionAction string
type RoleType string
type UserStatus string
type UserType string

const (
	// Permission Resources
	ResourceUser       PermissionResource = "user"
	ResourceRole       PermissionResource = "role"
	ResourcePermission PermissionResource = "permission"
	ResourceAuth       PermissionResource = "auth"
	ResourceToken      PermissionResource = "token"

	// Permission Actions
	ActionCreate     PermissionAction = "create"
	ActionRead       PermissionAction = "read"
	ActionUpdate     PermissionAction = "update"
	ActionDelete     PermissionAction = "delete"
	ActionAssign     PermissionAction = "assign"
	ActionRemove     PermissionAction = "remove"
	ActionActivate   PermissionAction = "activate"
	ActionDeactivate PermissionAction = "deactivate"

	// Role Types
	RoleTypeAdmin RoleType = "admin"
	RoleTypeUser  RoleType = "user"

	// User Status
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusLocked   UserStatus = "locked"

	// User Types
	UserTypeAdmin UserType = "admin"
	UserTypeUser  UserType = "user"
)