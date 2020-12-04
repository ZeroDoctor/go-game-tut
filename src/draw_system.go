package main

import (
	"github.com/faiface/pixel"
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/src/ecs"
)

type RenderComp struct {
	sprite *pixel.Sprite
	batch  *pixel.Batch
}

func (r RenderComp) Name() string {
	return "render"
}

type DrawSystem struct {
	*ecs.System
}

func NewDrawSystem() *DrawSystem {
	return &DrawSystem{
		System: ecs.NewSystem("render", "position"),
	}
}

func (d *DrawSystem) Update(dt float64) {

	d.Entities(func(iterator rbt.RbIterator, key rbt.RbKey, value interface{}) {
		render := d.InspectEntity(value).GetComponent("render").(*RenderComp)
		pos := d.InspectEntity(value).GetComponent("position").(*PositionComp)

		render.sprite.Draw(render.batch, pixel.IM.Moved(pixel.Vec{X: pos.x, Y: pos.y}))
	})
}
