package GoUtils

import "strconv"

// IsNumeric Checks if a String contains a number
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
