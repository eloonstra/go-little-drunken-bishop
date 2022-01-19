package drunkenbishop

import (
	"github.com/eloonstra/go-little-drunken-bishop/internal/model/board"
	"github.com/eloonstra/go-little-drunken-bishop/internal/model/heatmap"
)

func GenerateRandomArt(width, height int, bytes []byte, includeBorders bool, title string) string {
	return heatmap.New(board.New(width, height), bytes).ToString(includeBorders, title)
}

func GenerateHeatmap(width, height int, bytes []byte, includeBorders bool, title string) [][]int {
	return heatmap.New(board.New(width, height), bytes).Heatmap()
}
