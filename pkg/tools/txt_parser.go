package tools

import (
	"os"
	"strings"
)

func GetTxtValues(path string) (map[string]string, error) {
	var mappedValues = map[string]string{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line) // No whitespaces
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Ignore comments
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Ignore lines without "="
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		mappedValues[key] = value
	}

	return mappedValues, nil
}
