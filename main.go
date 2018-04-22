package main

import (
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type point struct {
	x, y int
}

const coldef = termbox.ColorDefault

func setupInput() {
	termbox.SetInputMode(termbox.InputAlt)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	setupInput()

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
