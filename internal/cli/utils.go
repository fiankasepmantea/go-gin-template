package cli

import "strings"

func toPascalCase(s string) string {

	parts := strings.Split(s, "_")

	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}