package asciiart

import "strings"

func ContainsOnly(char string) bool {
	for i := 0; i < len(char); i++ {
		if !strings.ContainsAny(string(char[i]), "\\n") {
			return false
		}
	}
	return true
}
