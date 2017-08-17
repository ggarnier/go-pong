package main

import (
	"math/rand"
	"time"

	tm "github.com/buger/goterm"
)

type point struct {
	x, y int
}

func main() {
	maxSpeed := point{
		x: 6,
		y: 2,
	}
	position := point{
		x: 40 | tm.PCT,
		y: 40 | tm.PCT,
	}
	rand.Seed(time.Now().UnixNano())
	box := tm.NewBox(2, 2, 0)

	for {
		tm.Clear()
		tm.MoveCursor(1, 1)

		position.x += rand.Intn(2*maxSpeed.x) - maxSpeed.x
		position.y += rand.Intn(2*maxSpeed.y) - maxSpeed.y
		tm.Print(tm.MoveTo(box.String(), position.x, position.y))

		tm.Flush()

		time.Sleep(300 * time.Millisecond)
	}
}
