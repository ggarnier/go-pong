package main

import "math/rand"

type ball struct {
	position point
	maxSpeed point
	speed    point
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
