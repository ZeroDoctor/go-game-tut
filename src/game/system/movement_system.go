package system

import (
	"github.com/ByteArena/box2d"
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/ecs"
	"github.com/zerodoctor/go-tut/game/comp"
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
		entity := m.InspectEntity(value)
		pos := entity.GetComponent("position").(*comp.PositionComp)
		vel := entity.GetComponent("velocity").(*comp.VelocityComp)

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

		boxVel := pos.Body.GetLinearVelocity()

		boxVelChangeX := vel.Vx - boxVel.X
		boxVelChangeY := vel.Vy - boxVel.Y
		impulseX := pos.Body.GetMass() * boxVelChangeX
		impulseY := pos.Body.GetMass() * boxVelChangeY

		velocity := box2d.B2Vec2{X: impulseX, Y: impulseY}
		position := pos.Body.GetPosition()

		pos.Body.ApplyLinearImpulse(velocity, position, true)
	})
}
