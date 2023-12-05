package helper

func StringInArray(target string, arr []string) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}
