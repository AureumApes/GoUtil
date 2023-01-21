package GoUtils

// PointerStringToString removes the pointer from a string
func PointerStringToString(str *string) string {
	return *str
}
