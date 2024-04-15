package sudoku_board

import (
	"ads-cw/pkg/display"
	"fmt"
	"strings"
)

const (
	verticalDividerSymbol     = "-"
	horizontalDividerSymbol   = "|"
	crossSectionDividerSymbol = "+"
	noValueSymbol             = "."
	resetStyle                = "\033[0m"          // anything after this will have no custom colouring applied
	whiteBGBlack              = "\033[47m\033[30m" // White background, Black text
	invalidCell               = "\033[41m\033[37m"
	invalidSelectedCell       = "\033[41m\033[30m"
	staticCell                = "\033[100m\033[30m"
)

func (b *Board) serializeBoard(pointer *display.Pointer) string {
	var sb strings.Builder

	for y, row := range b.Content {
		if y%b.subGridSize == 0 && y != 0 {
			sb.WriteString(createVerticalDivider(b.subGridSize))
		}
		for x, val := range row {
			if x%b.subGridSize == 0 && x != 0 {
				sb.WriteString(horizontalDividerSymbol + " ")
			}

			tileColour := ""
			if b.initialBoard[y][x] != 0 {
				tileColour = staticCell
			}
			if _, ok := b.invalidPlacements[[2]int{y, x}]; ok {
				tileColour = invalidCell
			}
			if x == pointer.X && y == pointer.Y {
				if tileColour == invalidCell {
					tileColour = invalidSelectedCell
				} else {
					tileColour = pointer.SelectedTileColour
				}
			}

			func() {
				defer sb.WriteString(" ")
				if tileColour != "" {
					defer sb.WriteString(resetStyle)
					sb.WriteString(tileColour)
				}

				if val == 0 {
					sb.WriteString(noValueSymbol)
				} else {
					sb.WriteString(fmt.Sprintf("%d", val))
				}
			}()
		}
		sb.WriteString("\n")
	}

	sb.WriteString(invalidCell)
	sb.WriteString(b.footerMessage)
	sb.WriteString(resetStyle)

	return sb.String()
}

func createVerticalDivider(size int) (divider string) {

	squareDivider := createLine(size * 2)
	divider = squareDivider

	for i := 0; i < size-2; i++ {
		divider += crossSectionDividerSymbol + squareDivider + verticalDividerSymbol
	}

	return divider + crossSectionDividerSymbol + squareDivider + "\n"
}

func createLine(length int) (line string) {
	for i := 0; i < length; i++ {
		line += verticalDividerSymbol
	}

	return line
}
