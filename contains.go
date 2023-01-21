package GoUtils

// Contains checks if a slice contains a index with a certain value
func Contains(slice []any, value any) bool {
	for _, curr := range slice {
		if curr == value {
			return true
		}
	}
	return false
}
