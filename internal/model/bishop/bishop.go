package bishop

type Bishop struct {
	startX, startY, endX, endY int
}

func New(startX, startY int) *Bishop {
	return &Bishop{
		startX: startX,
		startY: startY,
		endX:   startX,
		endY:   startY,
	}
}

func (bishop *Bishop) Move(x, y int) {
	bishop.endX, bishop.endY = x, y
}

func (bishop *Bishop) StartX() int {
	return bishop.startX
}

func (bishop *Bishop) StartY() int {
	return bishop.startY
}

func (bishop *Bishop) EndX() int {
	return bishop.endX
}

func (bishop *Bishop) EndY() int {
	return bishop.endY
}

func (bishop *Bishop) SetStartX(x int) {
	bishop.startX = x
}

func (bishop *Bishop) SetStartY(y int) {
	bishop.startY = y
}

func (bishop *Bishop) Reset() {
	bishop.endX, bishop.endY = bishop.startX, bishop.startY
}
