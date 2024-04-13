package display

import (
	"strings"
)

func draw(components [][]*ComponentNode) string {
	var builder strings.Builder

	for _, row := range components {
		var items []string
		for _, componentNode := range row {
			items = append(items, componentNode.Component.Serialize(componentNode.Pointer))
		}
		// Draw each row of components side by side, with a specified number of spaces in between
		builder.WriteString(SideBySide(items, 4))
		builder.WriteString("\n\n") // Add spacing between rows
	}

	return builder.String()
}
