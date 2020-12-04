package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/zerodoctor/go-tut/src/ecs"
	"github.com/zerodoctor/go-tut/src/util"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Your mom 2.0",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	engine := ecs.Engine{
		Win: win,
		Cfg: cfg,
	}

	hRender := &RenderComp{
		sprite: util.GetSprite("res/entity/wizard_player.png", 0, 0, 64, 64),
		batch:  util.GetBatch("res/entity/wizard_player.png"),
	}
	hPosition := &PositionComp{x: 512.0, y: 384.0}
	hVelocity := &VelocityComp{maxv: 0.95, minv: -0.95, speed: 0.9}
	hKeyDirection := &DirectionComp{}

	hero := &Hero{
		Entity: ecs.EntityFactory.NewEntity(
			"Player", hRender, hPosition,
			hVelocity, hKeyDirection,
		),
	}

	engine.AddSystem(NewKeySystem(win))
	engine.AddSystem(NewMovementSystem())
	engine.AddSystem(NewDrawSystem())
	engine.AddEntity(hero)
	engine.Update()

}

func main() {
	pixelgl.Run(run)
}
