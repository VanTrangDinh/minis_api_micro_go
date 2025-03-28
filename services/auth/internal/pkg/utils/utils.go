package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// ParseTime parses a time string in RFC3339 format
func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeStr)
}

// FormatTime formats a time.Time to RFC3339 string
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// IsExpired checks if a timestamp has expired
func IsExpired(timestamp time.Time) bool {
	return time.Now().After(timestamp)
}

// GetExpirationTime returns a time.Time that expires after the specified duration
func GetExpirationTime(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	// Implementation will be added when needed
	return "", nil
}

// ComparePasswords compares a password with its hash
func ComparePasswords(password, hash string) bool {
	// Implementation will be added when needed
	return false
}

// GenerateOTP generates a 6-digit OTP
func GenerateOTP() (string, error) {
	b := make([]byte, 3)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	otp := int(b[0])<<16 | int(b[1])<<8 | int(b[2])
	return fmt.Sprintf("%06d", otp%1000000), nil
}

// ValidateOTP validates a 6-digit OTP
func ValidateOTP(otp string) bool {
	if len(otp) != 6 {
		return false
	}
	for _, c := range otp {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
} 