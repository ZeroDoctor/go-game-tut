package system

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/src/ecs"
	"github.com/zerodoctor/go-tut/src/game/comp"
)

type KeySystem struct {
	name string
	Win  *pixelgl.Window
	*ecs.System
}

func NewKeySystem(win *pixelgl.Window) *KeySystem {
	return &KeySystem{
		name:   "KeySystem",
		Win:    win,
		System: ecs.NewSystem("velocity", "direction"),
	}
}

func (k *KeySystem) Name() string {
	return k.name
}

func (k *KeySystem) Update(dt float64) {
	k.Entities(func(iterator rbt.RbIterator, key rbt.RbKey, value interface{}) {
		vel := k.InspectEntity(value).GetComponent("velocity").(*comp.VelocityComp)
		dkey := k.InspectEntity(value).GetComponent("direction").(*comp.DirectionComp)

		if k.Win.Pressed(pixelgl.KeyA) {
			vel.Vx -= vel.Speed * dt
			dkey.Left = true
		} else if vel.Vx < 0.0 {
			vel.Vx += vel.Speed * dt * 1.875
			if vel.Vx > 0.0 {
				vel.Vx = 0.0
			}
			dkey.Left = false
		}

		if k.Win.Pressed(pixelgl.KeyD) {
			vel.Vx += vel.Speed * dt
			dkey.Right = true
		} else if vel.Vx > 0.0 {
			vel.Vx -= vel.Speed * dt * 1.875
			if vel.Vx < 0.0 {
				vel.Vx = 0.0
			}
			dkey.Right = false
		}

		if k.Win.Pressed(pixelgl.KeyS) {
			vel.Vy -= vel.Speed * dt
			dkey.Down = true
		} else if vel.Vy < 0.0 {
			vel.Vy += vel.Speed * dt * 1.875
			if vel.Vy > 0.0 {
				vel.Vy = 0.0
			}
			dkey.Down = false
		}

		if k.Win.Pressed(pixelgl.KeyW) {
			vel.Vy += vel.Speed * dt
			dkey.Up = true
		} else if vel.Vy > 0.0 {
			vel.Vy -= vel.Speed * dt * 1.875
			if vel.Vy < 0.0 {
				vel.Vy = 0.0
			}
			dkey.Up = false
		}
	})
}
