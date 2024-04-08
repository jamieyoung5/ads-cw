package display

import (
	"strings"
)

func draw(component *ComponentNode) string {
	var builder strings.Builder

	topLeftMostComponent := getTopLeftmostComponent(component)
	drawRow(topLeftMostComponent, &builder)
	for topLeftMostComponent.Down != nil {
		drawRow(topLeftMostComponent.Down, &builder)
		builder.WriteString("\n\n")
		topLeftMostComponent = topLeftMostComponent.Down
	}

	return builder.String()
}

func drawRow(component *ComponentNode, builder *strings.Builder) {
	items := []string{component.Component.Serialize(component.Pointer.x, component.Pointer.y, component.Pointer.selectedTileColour)}

	for component.Right != nil {
		items = append(items, component.Right.Component.Serialize(component.Right.Pointer.x, component.Right.Pointer.y, component.Right.Pointer.selectedTileColour))
		component = component.Right
	}

	builder.WriteString(SideBySide(items, 4))
}

func getTopLeftmostComponent(component *ComponentNode) *ComponentNode {
	for component.Up != nil {
		component = component.Up
	}

	for component.Left != nil {
		component = component.Left
	}

	return component
}
