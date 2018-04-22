package main

import termbox "github.com/nsf/termbox-go"

type screen struct {
	ball *ball
	size point
}

const SCREEN_SIZE_X = 100
const SCREEN_SIZE_Y = 20

func (s *screen) Clear() {
	termbox.Clear(coldef, coldef)
}

func (s *screen) Render() {
	s.drawBorder()
	tbprint(s.ball.position.x, s.ball.position.y, coldef, coldef, s.ball.String())
	termbox.Flush()
}

func (s *screen) drawBorder() {
	hborder := ""
	for i := 0; i < s.size.x; i++ {
		hborder += "-"
	}
	tbprint(0, 0, coldef, coldef, hborder)
	tbprint(0, s.size.y-1, coldef, coldef, hborder)
	for i := 1; i < s.size.y-1; i++ {
		tbprint(0, i, coldef, coldef, "|")
		tbprint(s.size.x-1, i, coldef, coldef, "|")
	}
}

func NewScreen(b *ball) *screen {
	return &screen{
		ball: b,
		size: point{
			x: SCREEN_SIZE_X,
			y: SCREEN_SIZE_Y,
		},
	}
}
