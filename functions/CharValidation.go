package functions

func charValidation(str string) bool {
	slice := []rune(str)
	for _, char := range slice {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
