package response

// Code represents the response code
type Code int

const (
	// Success codes (1xxx)
	Success Code = 1000
	Created Code = 1001
	Updated Code = 1002
	Deleted Code = 1003

	// Client error codes (2xxx)
	BadRequest         Code = 2000
	Unauthorized       Code = 2001
	Forbidden          Code = 2002
	NotFound           Code = 2003
	ValidationError    Code = 2004
	DuplicateEntry     Code = 2005
	InvalidCredentials Code = 2006
	AccountDisabled    Code = 2007
	TokenExpired       Code = 2008
	TokenInvalid       Code = 2009

	// Server error codes (3xxx)
	InternalServerError Code = 3000
	DatabaseError       Code = 3001
	CacheError          Code = 3002
	ServiceUnavailable  Code = 3003
)

var codeMessages = map[Code]string{
	// Success messages
	Success: "Success",
	Created: "Resource created successfully",
	Updated: "Resource updated successfully",
	Deleted: "Resource deleted successfully",

	// Client error messages
	BadRequest:         "Bad request",
	Unauthorized:       "Unauthorized",
	Forbidden:          "Forbidden",
	NotFound:           "Resource not found",
	ValidationError:    "Validation error",
	DuplicateEntry:     "Duplicate entry",
	InvalidCredentials: "Invalid credentials",
	AccountDisabled:    "Account is disabled",
	TokenExpired:       "Token has expired",
	TokenInvalid:       "Invalid token",

	// Server error messages
	InternalServerError: "Internal server error",
	DatabaseError:       "Database error",
	CacheError:          "Cache error",
	ServiceUnavailable:  "Service unavailable",
}

// Message returns the default message for each code
func (c Code) Message() string {
	if msg, ok := codeMessages[c]; ok {
		return msg
	}
	return "Unknown error"
}
