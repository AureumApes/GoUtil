package GoUtils

func Contains(slice []any, value any) bool {
	for _, curr := range slice {
		if curr == value {
			return true
		}
	}
	return false
}
