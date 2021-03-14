package comp

import (
	"github.com/ByteArena/box2d"
	"github.com/faiface/pixel"
)

type PositionComp struct {
	X, Y float64
	Body *box2d.B2Body
}

func (p PositionComp) Name() string {
	return "position"
}

type VelocityComp struct {
	Vx, Vy     float64
	Maxv, Minv float64
	Speed      float64
	DeAccel    float64
}

func (v VelocityComp) Name() string {
	return "velocity"
}

type CameraComp struct {
	Zoom      float64
	ZoomSpeed float64
	Speed     float64
	X, Y      float64
}

func (c CameraComp) Name() string {
	return "camera"
}

type RenderComp struct {
	Sprite *pixel.Sprite
	Batch  *pixel.Batch
	Scale  float64
}

func (r RenderComp) Name() string {
	return "render"
}

type DirectionComp struct {
	Up, Down, Left, Right bool
}

func (k DirectionComp) Name() string {
	return "direction"
}
