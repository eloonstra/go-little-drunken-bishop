package board

import (
	"github.com/eloonstra/go-little-drunken-bishop/internal/model/bishop"
	"math"
)

type Board struct {
	width, height int
	bishop        *bishop.Bishop
}

func New(width, height int) *Board {
	return &Board{
		width:  width,
		height: height,
		bishop: bishop.New(width/2, height/2),
	}
}

func (board *Board) Height() int {
	return board.height
}

func (board *Board) Width() int {
	return board.width
}

func (board *Board) Bishop() *bishop.Bishop {
	return board.bishop
}

func (board *Board) MoveBishop(x, y int) {
	board.Bishop().Move(
		int(math.Max(0, math.Min(float64(x), float64(board.width-1)))),
		int(math.Max(0, math.Min(float64(y), float64(board.height-1)))),
	)
}
