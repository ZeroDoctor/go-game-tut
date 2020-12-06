package system

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/src/ecs"
	"github.com/zerodoctor/go-tut/src/game/comp"
)

type CamSystem struct {
	name string
	Win  *pixelgl.Window
	*ecs.System
}

func NewCamSystem(win *pixelgl.Window) *CamSystem {
	return &CamSystem{
		name:   "CamSystem",
		Win:    win,
		System: ecs.NewSystem("position", "camera"),
	}
}

func (c *CamSystem) Name() string {
	return c.name
}

func (c *CamSystem) Update(dt float64) {
	c.Entities(func(iterator rbt.RbIterator, key rbt.RbKey, value interface{}) {

		camera := c.InspectEntity(value).GetComponent("camera").(*comp.CameraComp)
		targetPos := c.InspectEntity(value).GetComponent("position").(*comp.PositionComp)

		camera.X = targetPos.X
		camera.Y = targetPos.Y

		camPos := pixel.Vec{X: camera.X, Y: camera.Y}

		cam := pixel.IM.Scaled(camPos, camera.Zoom).Moved(c.Win.Bounds().Center().Sub(camPos))
		c.Win.SetMatrix(cam)
	})
}
