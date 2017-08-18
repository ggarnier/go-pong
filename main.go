package main

import (
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type screen struct {
	ball *ball
	size point
}

type ball struct {
	position point
	maxSpeed point
	speed    point
}

type point struct {
	x, y int
}

const coldef = termbox.ColorDefault

func (s *screen) Clear() {
	termbox.Clear(coldef, coldef)
}

func (s *screen) Render() {
	tbprint(s.ball.position.x, s.ball.position.y, coldef, coldef, s.ball.String())
	termbox.Flush()
}

func (b *ball) Move() {
	b.position.x += b.speed.x
	b.position.y += b.speed.y
	if b.position.x <= 0 || b.position.x >= 100 {
		b.speed.x *= -1
	}
	if b.position.y <= 0 || b.position.y >= 20 {
		b.speed.y *= -1
	}
}

func (b *ball) String() string {
	return "o"
}

func NewScreen(b *ball) *screen {
	return &screen{
		ball: b,
		size: point{
			x: 100,
			y: 20,
		},
	}
}

func NewBall() *ball {
	startingPoint := point{
		x: 50,
		y: 10,
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
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)

	time.Sleep(2 * time.Second)

	ball := NewBall()
	screen := NewScreen(ball)

	for i := 0; i < 30; i++ {
		screen.Clear()

		ball.Move()
		screen.Render()

		time.Sleep(300 * time.Millisecond)
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
