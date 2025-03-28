package handler

import (
	"net/http"
	"strconv"

	"minisapi/services/auth/internal/common"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RoleHandler handles role-related requests
type RoleHandler struct {
	roleRepo       repository.RoleRepository
	permissionRepo repository.PermissionRepository
}

// NewRoleHandler creates a new instance of RoleHandler
func NewRoleHandler(roleRepo repository.RoleRepository, permissionRepo repository.PermissionRepository) *RoleHandler {
	return &RoleHandler{
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
	}
}

// CreateRoleRequest represents the create role request body
type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required" example:"admin"`
	Description string `json:"description" example:"Administrator role with full access"`
	Permissions []uint `json:"permissions" example:"[1,2,3]"`
}

// Create godoc
// @Summary Create a new role
// @Description Create a new role with optional permissions
// @Tags roles
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body CreateRoleRequest true "Role details"
// @Success 201 {object} entity.Role
// @Failure 400 {object} entity.ErrorResponse
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 500 {object} entity.ErrorResponse
// @Router /api/roles [post]
func (h *RoleHandler) Create(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.VALIDATION_ERROR)
		return
	}

	// Check if role name already exists
	if _, err := h.roleRepo.FindByName(req.Name); err == nil {
		c.JSON(http.StatusBadRequest, common.ROLE_EXISTS)
		return
	}

	// Create role
	role := &entity.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	// Add permissions if provided
	if len(req.Permissions) > 0 {
		permissions := make([]entity.Permission, 0, len(req.Permissions))
		for _, permID := range req.Permissions {
			perm, err := h.permissionRepo.FindByID(permID)
			if err != nil {
				c.JSON(http.StatusBadRequest, common.INVALID_PERMISSION)
				return
			}
			permissions = append(permissions, *perm)
		}
		role.Permissions = permissions
	}

	if err := h.roleRepo.Create(role); err != nil {
		c.JSON(http.StatusInternalServerError, common.DB_ERROR)
		return
	}

	c.JSON(http.StatusCreated, role)
}

// List godoc
// @Summary List all roles
// @Description Get a paginated list of roles
// @Tags roles
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} entity.Role
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 500 {object} entity.ErrorResponse
// @Router /api/roles [get]
func (h *RoleHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	roles, total, err := h.roleRepo.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DB_ERROR)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  roles,
		"total": total,
	})
}

// Get godoc
// @Summary Get a role by ID
// @Description Get detailed information about a role
// @Tags roles
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Role ID"
// @Success 200 {object} entity.Role
// @Failure 400 {object} entity.ErrorResponse
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 404 {object} entity.ErrorResponse
// @Router /api/roles/{id} [get]
func (h *RoleHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.PARAMS_ERROR)
		return
	}

	role, err := h.roleRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.ROLE_NOT_FOUND)
		return
	}

	c.JSON(http.StatusOK, role)
}

// Update godoc
// @Summary Update a role
// @Description Update a role's details and permissions
// @Tags roles
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Role ID"
// @Param request body CreateRoleRequest true "Role details"
// @Success 200 {object} entity.Role
// @Failure 400 {object} entity.ErrorResponse
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 404 {object} entity.ErrorResponse
// @Failure 500 {object} entity.ErrorResponse
// @Router /api/roles/{id} [put]
func (h *RoleHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.PARAMS_ERROR)
		return
	}

	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.VALIDATION_ERROR)
		return
	}

	role, err := h.roleRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.ROLE_NOT_FOUND)
		return
	}

	role.Name = req.Name
	role.Description = req.Description

	// Update permissions if provided
	if len(req.Permissions) > 0 {
		permissions := make([]entity.Permission, 0, len(req.Permissions))
		for _, permID := range req.Permissions {
			perm, err := h.permissionRepo.FindByID(permID)
			if err != nil {
				c.JSON(http.StatusBadRequest, common.INVALID_PERMISSION)
				return
			}
			permissions = append(permissions, *perm)
		}
		role.Permissions = permissions
	}

	if err := h.roleRepo.Update(role); err != nil {
		c.JSON(http.StatusInternalServerError, common.DB_ERROR)
		return
	}

	c.JSON(http.StatusOK, role)
}

// Delete godoc
// @Summary Delete a role
// @Description Delete a role by ID
// @Tags roles
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Role ID"
// @Success 200 {object} entity.MessageResponse
// @Failure 400 {object} entity.ErrorResponse
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 500 {object} entity.ErrorResponse
// @Router /api/roles/{id} [delete]
func (h *RoleHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.PARAMS_ERROR)
		return
	}

	if err := h.roleRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, common.DB_ERROR)
		return
	}

	c.JSON(http.StatusOK, common.SUCCESS)
}
