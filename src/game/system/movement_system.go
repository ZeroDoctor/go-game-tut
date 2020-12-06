package system

import (
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/src/ecs"
	"github.com/zerodoctor/go-tut/src/game/comp"
)

type MovementSystem struct {
	name string
	*ecs.System
}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{
		name: "MovementSystem",
		System: ecs.NewSystem(
			"position", "velocity",
		),
	}
}

func (m *MovementSystem) Name() string {
	return m.name
}

func (m *MovementSystem) Update(dt float64) {
	m.Entities(func(iterator rbt.RbIterator, key rbt.RbKey, value interface{}) {
		pos := m.InspectEntity(value).GetComponent("position").(*comp.PositionComp)
		vel := m.InspectEntity(value).GetComponent("velocity").(*comp.VelocityComp)

		if vel.Vx > vel.Maxv {
			vel.Vx = vel.Maxv
		}
		if vel.Vy > vel.Maxv {
			vel.Vy = vel.Maxv
		}

		if vel.Vx < vel.Minv {
			vel.Vx = vel.Minv
		}
		if vel.Vy < vel.Minv {
			vel.Vy = vel.Minv
		}

		pos.X += vel.Vx
		pos.Y += vel.Vy
	})
}
