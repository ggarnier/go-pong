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
		if position.x <= 0 {
			position.x = 1
		}
		if position.x >= 100|tm.PCT {
			position.x = 100 | tm.PCT - 1
		}
		position.y += rand.Intn(2*maxSpeed.y) - maxSpeed.y
		if position.y <= 0 {
			position.y = 1
		}
		if position.y >= 100|tm.PCT {
			position.y = 100 | tm.PCT - 1
		}
		tm.Print(tm.MoveTo(box.String(), position.x, position.y))

		tm.Flush()

		time.Sleep(300 * time.Millisecond)
	}
}
