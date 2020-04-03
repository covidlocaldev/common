package utils

func ContainsString(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}

	return false
}