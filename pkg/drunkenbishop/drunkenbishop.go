package drunkenbishop

import (
	"github.com/eloonstra/go-little-drunken-bishop/internal/model/board"
	"github.com/eloonstra/go-little-drunken-bishop/internal/model/heatmap"
)

// GenerateRandomArt generates a random art using the drunken bishop algorithm.
func GenerateRandomArt(width, height int, bytes []byte, includeBorders bool, title string) string {
	return heatmap.New(board.New(width, height), bytes).ToString(includeBorders, title)
}

// GenerateHeatmap generates a heatmap using the drunken bishop algorithm.
func GenerateHeatmap(width, height int, bytes []byte) [][]int {
	return heatmap.New(board.New(width, height), bytes).Heatmap()
}
