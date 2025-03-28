// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/permissions": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a paginated list of permissions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "permissions"
                ],
                "summary": "List all permissions",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Permission"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new permission with the specified details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "permissions"
                ],
                "summary": "Create a new permission",
                "parameters": [
                    {
                        "description": "Permission details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreatePermissionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Permission"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/permissions/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get detailed information about a permission",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "permissions"
                ],
                "summary": "Get a permission by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Permission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Permission"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update a permission's details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "permissions"
                ],
                "summary": "Update a permission",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Permission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Permission details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreatePermissionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Permission"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a permission by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "permissions"
                ],
                "summary": "Delete a permission",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Permission ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/roles": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a paginated list of roles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "List all roles",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Role"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new role with optional permissions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Create a new role",
                "parameters": [
                    {
                        "description": "Role details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateRoleRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/roles/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get detailed information about a role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Get a role by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update a role's details and permissions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Update a role",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Role details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateRoleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Role"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a role by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Delete a role",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Role ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate user and return JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Registration details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.AuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "enums.PermissionAction": {
            "type": "string",
            "enum": [
                "create",
                "read",
                "update",
                "delete",
                "assign",
                "remove",
                "activate",
                "deactivate"
            ],
            "x-enum-varnames": [
                "ActionCreate",
                "ActionRead",
                "ActionUpdate",
                "ActionDelete",
                "ActionAssign",
                "ActionRemove",
                "ActionActivate",
                "ActionDeactivate"
            ]
        },
        "enums.PermissionResource": {
            "type": "string",
            "enum": [
                "user",
                "role",
                "permission",
                "auth",
                "token"
            ],
            "x-enum-varnames": [
                "ResourceUser",
                "ResourceRole",
                "ResourcePermission",
                "ResourceAuth",
                "ResourceToken"
            ]
        },
        "enums.RoleType": {
            "type": "string",
            "enum": [
                "system",
                "admin",
                "user",
                "manager",
                "editor",
                "viewer"
            ],
            "x-enum-varnames": [
                "RoleTypeSystem",
                "RoleTypeAdmin",
                "RoleTypeUser",
                "RoleTypeManager",
                "RoleTypeEditor",
                "RoleTypeViewer"
            ]
        },
        "enums.UserStatus": {
            "type": "string",
            "enum": [
                "active",
                "inactive",
                "blocked",
                "deleted"
            ],
            "x-enum-varnames": [
                "UserStatusActive",
                "UserStatusInactive",
                "UserStatusBlocked",
                "UserStatusDeleted"
            ]
        },
        "enums.UserType": {
            "type": "string",
            "enum": [
                "system",
                "admin",
                "user"
            ],
            "x-enum-varnames": [
                "UserTypeSystem",
                "UserTypeAdmin",
                "UserTypeUser"
            ]
        },
        "handlers.AuthResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "handlers.CreatePermissionRequest": {
            "type": "object",
            "required": [
                "action",
                "name",
                "resource"
            ],
            "properties": {
                "action": {
                    "type": "string",
                    "example": "create"
                },
                "description": {
                    "type": "string",
                    "example": "Permission to create new users"
                },
                "name": {
                    "type": "string",
                    "example": "create_user"
                },
                "resource": {
                    "type": "string",
                    "example": "users"
                }
            }
        },
        "handlers.CreateRoleRequest": {
            "type": "object"
        },
        "handlers.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "secret123"
                },
                "username": {
                    "type": "string",
                    "example": "johndoe"
                }
            }
        },
        "handlers.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "password": {
                    "type": "string",
                    "minLength": 6,
                    "example": "secret123"
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3,
                    "example": "johndoe"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "models.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "operation successful"
                }
            }
        },
        "models.Permission": {
            "type": "object",
            "properties": {
                "action": {
                    "$ref": "#/definitions/enums.PermissionAction"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "resource": {
                    "$ref": "#/definitions/enums.PermissionResource"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Role"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Role": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Permission"
                    }
                },
                "type": {
                    "$ref": "#/definitions/enums.RoleType"
                },
                "updated_at": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Role"
                    }
                },
                "status": {
                    "$ref": "#/definitions/enums.UserStatus"
                },
                "type": {
                    "$ref": "#/definitions/enums.UserType"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "JWT token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MinisAPI Auth Service",
	Description:      "A JWT-based authentication service with RBAC support.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
