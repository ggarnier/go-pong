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

func (b *ball) Move() {
	b.position.x += b.speed.x
	b.position.y += b.speed.y
	if b.position.x <= 1 || b.position.x >= SCREEN_SIZE_X-2 {
		b.speed.x *= -1
	}
	if b.position.y <= 1 || b.position.y >= SCREEN_SIZE_Y-2 {
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
			x: SCREEN_SIZE_X,
			y: SCREEN_SIZE_Y,
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
