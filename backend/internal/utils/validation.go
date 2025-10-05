package utils

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail checks if an email is valid
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// ValidatePassword checks if a password meets minimum requirements
func ValidatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long"
	}

	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		return false, "Password must contain at least one number"
	}

	return true, ""
}

// ValidateLanguageCode checks if a language code is valid ISO 639-1
func ValidateLanguageCode(code string) bool {
	// Basic validation - extend with full ISO 639-1 list if needed
	validCodes := []string{"en", "zh", "zh-HK", "zh-TW", "zh-CN", "es", "fr", "de", "ja", "ko"}
	code = strings.ToLower(code)

	for _, valid := range validCodes {
		if strings.ToLower(valid) == code {
			return true
		}
	}
	return false
}

// ValidateRole checks if a role is valid
func ValidateRole(role string) bool {
	validRoles := []string{"admin", "learner"}
	for _, valid := range validRoles {
		if valid == role {
			return true
		}
	}
	return false
}
