package display

import (
	"fmt"
	"strings"
	"time"
)

func (c *Canvas) draw(quit chan bool) {
	state := c.serialize()
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(1 * time.Second)
			newState := c.serialize()
			if newState != state {
				fmt.Print("\033[H\033[2J\033[3J")
				fmt.Printf(newState)
			}
		}
	}
}

func (c *Canvas) serialize() string {
	var builder strings.Builder
	components := c.States.Peek().Components

	for _, row := range components {
		var items []string
		for _, componentNode := range row {
			items = append(items, componentNode.Component.Serialize(componentNode.Pointer))
		}
		// Draw each row of Components side by side, with a specified number of spaces in between
		builder.WriteString(SideBySide(items, 4))
		builder.WriteString("\n\n") // Add spacing between rows
	}

	return builder.String()
}
