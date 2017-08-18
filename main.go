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
	speed    point
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
	tm.Print(tm.MoveTo(s.ball.String(), s.ball.position.x, s.ball.position.y))
	tm.Flush()
}

func newRandomPosition(reference, maxSpeed point) point {
	newPosition := point{
		x: reference.x + rand.Intn(2*maxSpeed.x) - maxSpeed.x,
		y: reference.y + rand.Intn(2*maxSpeed.y) - maxSpeed.y,
	}
	if newPosition.x <= 0 {
		newPosition.x = 1
	}
	if newPosition.x >= 100|tm.PCT {
		newPosition.x = 100 | tm.PCT - 1
	}
	if newPosition.y <= 0 {
		newPosition.y = 1
	}
	if newPosition.y >= 100|tm.PCT {
		newPosition.y = 100 | tm.PCT - 1
	}
	return newPosition
}

func (b *ball) Move() {
	b.position.x += b.speed.x
	b.position.y += b.speed.y
}

func (b *ball) String() string {
	return "o"
}

func NewScreen(b *ball) *screen {
	return &screen{
		ball: b,
	}
}

func NewBall() *ball {
	startingPoint := point{
		x: 40 | tm.PCT,
		y: 40 | tm.PCT,
	}
	maxSpeed := point{
		x: 4,
		y: 2,
	}
	speed := point{
		x: rand.Intn(maxSpeed.x+1) + 1,
		y: rand.Intn(maxSpeed.y+1) + 1,
	}
	return &ball{
		maxSpeed: maxSpeed,
		position: startingPoint,
		speed:    speed,
		box:      tm.NewBox(2, 2, 0),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ball := NewBall()
	screen := NewScreen(ball)

	for {
		screen.Clear()

		ball.Move()
		screen.Paint()

		time.Sleep(300 * time.Millisecond)
	}
}
