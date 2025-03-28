package validator

import (
	"fmt"
	"minisapi/services/auth/internal/domain/entity"
	"minisapi/services/auth/internal/pkg/errors"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type ValidationBuilder struct {
	errors []*ValidationError
}

func New() *ValidationBuilder {
	return &ValidationBuilder{
		errors: make([]*ValidationError, 0),
	}
}

func (v *ValidationBuilder) ValidateEmail(field, email string) *ValidationBuilder {
	if email == "" {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: "email is required",
		})
		return v
	}

	if !emailRegex.MatchString(email) {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: "invalid email format",
		})
	}
	return v
}

func (v *ValidationBuilder) ValidatePassword(field, password string) *ValidationBuilder {
	if password == "" {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: "password is required",
		})
		return v
	}

	if len(password) < 8 {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: "password must be at least 8 characters",
		})
	}

	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: "password must contain at least one uppercase letter",
		})
	}

	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: "password must contain at least one lowercase letter",
		})
	}

	if !strings.ContainsAny(password, "0123456789") {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: "password must contain at least one number",
		})
	}

	if !strings.ContainsAny(password, "!@#$%^&*") {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: "password must contain at least one special character",
		})
	}

	return v
}

func (v *ValidationBuilder) ValidateRequired(field, value string) *ValidationBuilder {
	if value == "" {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s is required", field),
		})
	}
	return v
}

func (v *ValidationBuilder) ValidateMinLength(field, value string, min int) *ValidationBuilder {
	if len(value) < min {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be at least %d characters", field, min),
		})
	}
	return v
}

func (v *ValidationBuilder) ValidateMaxLength(field, value string, max int) *ValidationBuilder {
	if len(value) > max {
		v.errors = append(v.errors, &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must not exceed %d characters", field, max),
		})
	}
	return v
}

func (v *ValidationBuilder) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *ValidationBuilder) Errors() []*ValidationError {
	return v.errors
}

type Validator interface {
	Validate(i interface{}) error
}

type validatorImpl struct {
	validate *validator.Validate
}

func NewValidator() Validator {
	return &validatorImpl{
		validate: validator.New(),
	}
}

func (v *validatorImpl) Validate(i interface{}) error {
	if err := v.validate.Struct(i); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Convert validation errors to custom errors
			for _, e := range validationErrors {
				switch e.Field() {
				case "Username":
					return errors.ErrInvalidInput
				case "Email":
					return errors.ErrInvalidInput
				case "Password":
					return errors.ErrInvalidInput
				case "FirstName":
					return errors.ErrInvalidInput
				case "LastName":
					return errors.ErrInvalidInput
				case "Phone":
					return errors.ErrInvalidInput
				default:
					return errors.ErrValidation
				}
			}
		}
		return errors.ErrValidation
	}

	// Additional custom validations
	switch t := i.(type) {
	case *entity.User:
		return v.validateUser(t)
	case *entity.Role:
		return v.validateRole(t)
	case *entity.Permission:
		return v.validatePermission(t)
	default:
		return nil
	}
}

func (v *validatorImpl) validateUser(user *entity.User) error {
	// Add custom user validations here
	return nil
}

func (v *validatorImpl) validateRole(role *entity.Role) error {
	// Add custom role validations here
	return nil
}

func (v *validatorImpl) validatePermission(permission *entity.Permission) error {
	// Add custom permission validations here
	return nil
}
