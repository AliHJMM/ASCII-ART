package functions

// CharValidation checks if all characters in the input string are valid ASCII characters (printable characters).
func CharValidation(str string) bool {
	slice := []rune(str)
	for _, char := range slice {
		// Check if the character is outside the range of printable ASCII characters
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
