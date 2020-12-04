package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/src/ecs"
)

type DirectionComp struct {
	up, down, left, right bool
}

func (k DirectionComp) Name() string {
	return "direction"
}

type KeySystem struct {
	Win *pixelgl.Window
	*ecs.System
}

func NewKeySystem(win *pixelgl.Window) *KeySystem {
	return &KeySystem{
		Win:    win,
		System: ecs.NewSystem("velocity", "direction"),
	}
}

func (k *KeySystem) Update(dt float64) {
	k.Entities(func(iterator rbt.RbIterator, key rbt.RbKey, value interface{}) {
		vel := k.InspectEntity(value).GetComponent("velocity").(*VelocityComp)
		dkey := k.InspectEntity(value).GetComponent("direction").(*DirectionComp)

		if k.Win.Pressed(pixelgl.KeyA) {
			vel.vx -= vel.speed * dt
			dkey.left = true
		} else if vel.vx < 0.0 {
			vel.vx += vel.speed * dt * 1.875
			if vel.vx > 0.0 {
				vel.vx = 0.0
			}
			dkey.left = false
		}

		if k.Win.Pressed(pixelgl.KeyD) {
			vel.vx += vel.speed * dt
			dkey.right = true
		} else if vel.vx > 0.0 {
			vel.vx -= vel.speed * dt * 1.875
			if vel.vx < 0.0 {
				vel.vx = 0.0
			}
			dkey.right = false
		}

		if k.Win.Pressed(pixelgl.KeyS) {
			vel.vy -= vel.speed * dt
			dkey.down = true
		} else if vel.vy < 0.0 {
			vel.vy += vel.speed * dt * 1.875
			if vel.vy > 0.0 {
				vel.vy = 0.0
			}
			dkey.down = false
		}

		if k.Win.Pressed(pixelgl.KeyW) {
			vel.vy += vel.speed * dt
			dkey.up = true
		} else if vel.vy > 0.0 {
			vel.vy -= vel.speed * dt * 1.875
			if vel.vy < 0.0 {
				vel.vy = 0.0
			}
			dkey.up = false
		}
	})
}
