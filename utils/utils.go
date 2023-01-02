package utils

// OnlyDigits checks if a string contains only digits
func OnlyDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
