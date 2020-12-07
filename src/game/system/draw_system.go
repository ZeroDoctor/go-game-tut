package system

import (
	"github.com/faiface/pixel"
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/src/ecs"
	"github.com/zerodoctor/go-tut/src/game/comp"
	"github.com/zerodoctor/go-tut/src/util"
)

type DrawSystem struct {
	name string
	*ecs.System
}

func NewDrawSystem() *DrawSystem {
	return &DrawSystem{
		name:   "DrawSystem",
		System: ecs.NewSystem("render", "position"),
	}
}

func (d *DrawSystem) Name() string {
	return d.name
}

func (d *DrawSystem) Update(dt float64) {

	d.Entities(func(iterator rbt.RbIterator, key rbt.RbKey, value interface{}) {
		render := d.InspectEntity(value).GetComponent("render").(*comp.RenderComp)
		pos := d.InspectEntity(value).GetComponent("position").(*comp.PositionComp)

		if render.Scale == 0.0 {
			render.Scale = 1.0
		}

		bodyPos := pos.Body.GetPosition()
		x := util.MetersToPixel(bodyPos.X)
		y := util.MetersToPixel(bodyPos.Y)
		render.Sprite.Draw(render.Batch, pixel.IM.Scaled(pixel.ZV, render.Scale).Moved(pixel.Vec{X: x, Y: y}))
	})
}
