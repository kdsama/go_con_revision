package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

// create all the dots , and put them on the map
func createBoid(bid int) {
	b := Boid{position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		// velocity between 1 and -1
		velocity: Vector2D{x: (rand.Float64() * 2) - 1, y: (rand.Float64() * 2) - 1},
		id:       bid,
	}
	boids[bid] = &b

	go b.start()
}

func (b *Boid) start() {

	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{-b.velocity.x, b.velocity.y}
	}
	if next.y >= screenWidth || next.y < 0 {
		b.velocity = Vector2D{-b.velocity.x, -b.velocity.y}
	}
}