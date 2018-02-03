package postal

import "strings"

func nonEmpty(input []string) []string {
	output := []string{}
	for _, part := range input {
		if strings.TrimSpace(part) != "" {
			output = append(output, strings.TrimSpace(part))
		}
	}
	return output
}

func first(input ...string) string {
	for _, part := range input {
		if strings.TrimSpace(part) != "" {
			return strings.TrimSpace(part)
		}
	}
	return ""
}
