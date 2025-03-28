package handler

import (
	"net/http"
	"strconv"

	"minisapi/services/auth/internal/common"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/enums"
	"minisapi/services/auth/internal/repository"

	"github.com/gin-gonic/gin"
)

// PermissionHandler handles permission-related requests
type PermissionHandler struct {
	permissionRepo repository.PermissionRepository
}

// NewPermissionHandler creates a new instance of PermissionHandler
func NewPermissionHandler(permissionRepo repository.PermissionRepository) *PermissionHandler {
	return &PermissionHandler{
		permissionRepo: permissionRepo,
	}
}

// CreatePermissionRequest represents the create permission request body
type CreatePermissionRequest struct {
	Name        string `json:"name" binding:"required" example:"create_user"`
	Description string `json:"description" example:"Permission to create new users"`
	Resource    string `json:"resource" binding:"required" example:"users"`
	Action      string `json:"action" binding:"required" example:"create"`
}

// Create godoc
// @Summary Create a new permission
// @Description Create a new permission with the specified details
// @Tags permissions
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body CreatePermissionRequest true "Permission details"
// @Success 201 {object} entity.Permission
// @Failure 400 {object} entity.ErrorResponse
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 500 {object} entity.ErrorResponse
// @Router /api/permissions [post]
func (h *PermissionHandler) Create(c *gin.Context) {
	var req CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.VALIDATION_ERROR)
		return
	}

	// Check if permission name already exists
	if _, err := h.permissionRepo.FindByName(req.Name); err == nil {
		c.JSON(http.StatusBadRequest, common.PERMISSION_EXISTS)
		return
	}

	// Validate resource and action
	resource := enums.PermissionResource(req.Resource)
	action := enums.PermissionAction(req.Action)

	// Create permission
	permission := &entity.Permission{
		Name:        req.Name,
		Description: req.Description,
		Resource:    string(resource),
		Action:      string(action),
	}

	if err := h.permissionRepo.Create(permission); err != nil {
		c.JSON(http.StatusInternalServerError, common.DB_ERROR)
		return
	}

	c.JSON(http.StatusCreated, permission)
}

// List godoc
// @Summary List all permissions
// @Description Get a paginated list of permissions
// @Tags permissions
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} entity.Permission
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 500 {object} entity.ErrorResponse
// @Router /api/permissions [get]
func (h *PermissionHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	permissions, total, err := h.permissionRepo.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DB_ERROR)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  permissions,
		"total": total,
	})
}

// Get godoc
// @Summary Get a permission by ID
// @Description Get detailed information about a permission
// @Tags permissions
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Permission ID"
// @Success 200 {object} entity.Permission
// @Failure 400 {object} entity.ErrorResponse
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 404 {object} entity.ErrorResponse
// @Router /api/permissions/{id} [get]
func (h *PermissionHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.PARAMS_ERROR)
		return
	}

	permission, err := h.permissionRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.PERMISSION_NOT_FOUND)
		return
	}

	c.JSON(http.StatusOK, permission)
}

// Update godoc
// @Summary Update a permission
// @Description Update a permission's details
// @Tags permissions
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Permission ID"
// @Param request body CreatePermissionRequest true "Permission details"
// @Success 200 {object} entity.Permission
// @Failure 400 {object} entity.ErrorResponse
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 404 {object} entity.ErrorResponse
// @Failure 500 {object} entity.ErrorResponse
// @Router /api/permissions/{id} [put]
func (h *PermissionHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.PARAMS_ERROR)
		return
	}

	var req CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.VALIDATION_ERROR)
		return
	}

	permission, err := h.permissionRepo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.PERMISSION_NOT_FOUND)
		return
	}

	// Validate resource and action
	resource := enums.PermissionResource(req.Resource)
	action := enums.PermissionAction(req.Action)

	permission.Name = req.Name
	permission.Description = req.Description
	permission.Resource = string(resource)
	permission.Action = string(action)

	if err := h.permissionRepo.Update(permission); err != nil {
		c.JSON(http.StatusInternalServerError, common.DB_ERROR)
		return
	}

	c.JSON(http.StatusOK, permission)
}

// Delete godoc
// @Summary Delete a permission
// @Description Delete a permission by ID
// @Tags permissions
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Permission ID"
// @Success 200 {object} entity.MessageResponse
// @Failure 400 {object} entity.ErrorResponse
// @Failure 401 {object} entity.ErrorResponse
// @Failure 403 {object} entity.ErrorResponse
// @Failure 500 {object} entity.ErrorResponse
// @Router /api/permissions/{id} [delete]
func (h *PermissionHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.PARAMS_ERROR)
		return
	}

	if err := h.permissionRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, common.DB_ERROR)
		return
	}

	c.JSON(http.StatusOK, common.SUCCESS)
}
