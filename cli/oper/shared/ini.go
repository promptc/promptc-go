package shared

import "strings"

func IniToMap(ini string) map[string]string {
	m := make(map[string]string)
	for _, line := range strings.Split(ini, "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		m[key] = value
	}
	return m
}
