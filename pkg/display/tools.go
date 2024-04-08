package display

import "strings"

func SideBySide(items []string, spacing int) (result string) {
	maxHeight := getMaxHeight(items)

	for i := 0; i < maxHeight; i++ {
		var rowBuilder strings.Builder

		for _, item := range items {
			lines := strings.Split(item, "\n")
			maxLength := getMaxLength(lines) // Use lines instead of the entire item

			if i < len(lines) {
				rowBuilder.WriteString(lines[i])
				rowBuilder.WriteString(getCustomSizedWhitespace(maxLength - len(lines[i]) + spacing))
			} else {
				rowBuilder.WriteString(getCustomSizedWhitespace(maxLength + spacing))
			}
		}

		rowBuilder.WriteString("\n")
		result += rowBuilder.String()
	}

	return result
}

func getCustomSizedWhitespace(size int) string {
	whitespace := ""
	for i := 0; i < size; i++ {
		whitespace += " "
	}

	return whitespace
}

func getMaxHeight(items []string) int {
	maxHeight := 0
	for _, item := range items {
		lines := strings.Split(item, "\n")
		if len(lines) > maxHeight {
			maxHeight = len(lines)
		}
	}

	return maxHeight
}

func getMaxLength(lines []string) int {
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	return maxLength
}
