package main

import (
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/src/ecs"
)

type PositionComp struct {
	x, y float64
}

func (p PositionComp) Name() string {
	return "position"
}

type VelocityComp struct {
	vx, vy     float64
	maxv, minv float64
	speed      float64
}

func (v VelocityComp) Name() string {
	return "velocity"
}

type MovementSystem struct {
	*ecs.System
}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{
		System: ecs.NewSystem(
			"position", "velocity",
		),
	}
}

func (m *MovementSystem) Update(dt float64) {
	m.Entities(func(iterator rbt.RbIterator, key rbt.RbKey, value interface{}) {
		pos := m.InspectEntity(value).GetComponent("position").(*PositionComp)
		vel := m.InspectEntity(value).GetComponent("velocity").(*VelocityComp)

		if vel.vx > vel.maxv {
			vel.vx = vel.maxv
		}
		if vel.vy > vel.maxv {
			vel.vy = vel.maxv
		}

		if vel.vx < vel.minv {
			vel.vx = vel.minv
		}
		if vel.vy < vel.minv {
			vel.vy = vel.minv
		}

		pos.x += vel.vx
		pos.y += vel.vy
	})
}
