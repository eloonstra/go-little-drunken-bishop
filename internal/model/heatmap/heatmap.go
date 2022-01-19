package heatmap

import (
	"github.com/eloonstra/go-little-drunken-bishop/internal/model/board"
	"strings"
)

type Heatmap struct {
	board *board.Board
	bytes []byte
}

const (
	northWest  = 0
	northEast  = 1
	southWest  = 2
	southEast  = 3
	characters = " .o+=*B0X@%&#/^"
)

func New(board *board.Board, bytes []byte) *Heatmap {
	return &Heatmap{
		board: board,
		bytes: bytes,
	}
}

func (heatmap *Heatmap) Heatmap() [][]int {
	result := make([][]int, heatmap.Board().Height())
	for i := 0; i < heatmap.Board().Height(); i++ {
		result[i] = make([]int, heatmap.Board().Width())
	}

	heatmap.Board().Bishop().Reset()

	for _, singleByte := range heatmap.bytes {
		for i := 0; i < 4; i++ {
			switch int(singleByte>>(i*2)) & 3 {
			case northWest:
				heatmap.Board().MoveBishop(
					heatmap.Board().Bishop().EndX()-1,
					heatmap.Board().Bishop().EndY()-1,
				)
			case northEast:
				heatmap.Board().MoveBishop(
					heatmap.Board().Bishop().EndX()+1,
					heatmap.Board().Bishop().EndY()-1,
				)
			case southWest:
				heatmap.Board().MoveBishop(
					heatmap.Board().Bishop().EndX()-1,
					heatmap.Board().Bishop().EndY()+1,
				)
			case southEast:
				heatmap.Board().MoveBishop(
					heatmap.Board().Bishop().EndX()+1,
					heatmap.Board().Bishop().EndY()+1,
				)
			}
			result[heatmap.Board().Bishop().EndY()][heatmap.Board().Bishop().EndX()]++
		}
	}

	return result
}

func (heatmap *Heatmap) Board() *board.Board {
	return heatmap.board
}

func (heatmap *Heatmap) Bytes() []byte {
	return heatmap.bytes
}

func (heatmap *Heatmap) ToString(includeBorders bool, title string) string {
	var stringBuilder strings.Builder
	generatedHeatmap := heatmap.Heatmap()

	if includeBorders {
		stringBuilder.WriteString(heatmap.generateTopBorder(title))
	}

	for y, row := range generatedHeatmap {
		if includeBorders {
			stringBuilder.WriteString("|")
		}
		for x := range row {
			stringBuilder.WriteString(heatmap.getCharacter(x, y))
		}
		if includeBorders {
			stringBuilder.WriteString("|\n")
		} else {
			stringBuilder.WriteString("\n")
		}
	}

	if includeBorders {
		stringBuilder.WriteString("+" + strings.Repeat("-", heatmap.Board().Width()) + "+\n")
	}

	return stringBuilder.String()
}

func (heatmap *Heatmap) updateTitle(title string) string {
	if len(title) > 0 {
		if len(title)&1 == heatmap.Board().Width()&1 {
			title = "[" + title + "]"
		} else {
			title = "[ " + title + "]"
		}
	}
	return title
}

func (heatmap *Heatmap) generateTopBorder(title string) string {
	title = heatmap.updateTitle(title)
	if len(title) == 0 {
		return "+" + strings.Repeat("-", heatmap.Board().Width()) + "+\n"
	} else {
		return "+" + strings.Repeat("-", heatmap.Board().Width()/2-len(title)/2) + title + strings.Repeat("-", heatmap.Board().Width()/2-len(title)/2) + "+\n"
	}
}

func (heatmap *Heatmap) getCharacter(x, y int) string {
	if x == heatmap.Board().Bishop().EndX() && y == heatmap.Board().Bishop().EndY() {
		return "E"
	} else if x == heatmap.Board().Bishop().StartX() && y == heatmap.Board().Bishop().StartY() {
		return "S"
	} else {
		return string(characters[heatmap.Heatmap()[y][x]])
	}
}
