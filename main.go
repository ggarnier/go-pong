package main

import (
	"math/rand"
	"time"

	tm "github.com/buger/goterm"
)

type screen struct {
	ball *ball
}

type ball struct {
	position point
	maxSpeed point
	box      *tm.Box
}

type point struct {
	x, y int
}

func (s *screen) Clear() {
	tm.Clear()
	tm.MoveCursor(1, 1)
}

func (s *screen) Paint() {
	tm.Print(tm.MoveTo(s.ball.box.String(), s.ball.position.x, s.ball.position.y))
	tm.Flush()
}

func (b *ball) Move() {
	b.position.x += rand.Intn(2*b.maxSpeed.x) - b.maxSpeed.x
	if b.position.x <= 0 {
		b.position.x = 1
	}
	if b.position.x >= 100|tm.PCT {
		b.position.x = 100 | tm.PCT - 1
	}
	b.position.y += rand.Intn(2*b.maxSpeed.y) - b.maxSpeed.y
	if b.position.y <= 0 {
		b.position.y = 1
	}
	if b.position.y >= 100|tm.PCT {
		b.position.y = 100 | tm.PCT - 1
	}
}

func NewScreen(b *ball) *screen {
	return &screen{
		ball: b,
	}
}

func NewBall() *ball {
	return &ball{
		maxSpeed: point{
			x: 6,
			y: 2,
		},
		position: point{
			x: 40 | tm.PCT,
			y: 40 | tm.PCT,
		},
		box: tm.NewBox(2, 2, 0),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	b := NewBall()
	s := NewScreen(b)

	for {
		s.Clear()

		b.Move()
		s.Paint()

		time.Sleep(300 * time.Millisecond)
	}
}
