package game

import (
	"github.com/ByteArena/box2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/zerodoctor/go-tut/src/ecs"
	"github.com/zerodoctor/go-tut/src/game/comp"
	"github.com/zerodoctor/go-tut/src/game/system"
	"github.com/zerodoctor/go-tut/src/util"
)

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Your mom 2.0",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	gravity := box2d.MakeB2Vec2(0.0, 0.0)
	world := box2d.MakeB2World(gravity)

	engine := ecs.NewEngine(win, cfg, &world)

	hRender := &comp.RenderComp{
		Sprite: util.GetSprite("res/entity/wizard_player.png", 64*0, 64*7, 64*1, 64*8),
		Batch:  util.GetBatch("res/entity/wizard_player.png"),
		Scale:  1.5,
	}

	hbodydef := box2d.NewB2BodyDef()
	hbodydef.Type = box2d.B2BodyType.B2_dynamicBody
	hbodydef.Position.Set(0.0, 0.0)
	hbody := world.CreateBody(hbodydef)

	size := util.PixelToMeters(32.0)
	hpoly := box2d.B2PolygonShape{}
	hpoly.SetAsBox(size, size)
	hbody.CreateFixture(&hpoly, 0.0)

	hFixDef := box2d.B2FixtureDef{}
	hFixDef.Shape = &hpoly
	hFixDef.Density = 1.0
	hFixDef.Friction = 0.3

	hbody.CreateFixtureFromDef(&hFixDef)

	hPosition := &comp.PositionComp{X: 0.0, Y: 0.0, Body: hbody}
	hVelocity := &comp.VelocityComp{Maxv: 0.95, Minv: -0.95, Speed: 0.9}
	hKeyDirection := &comp.DirectionComp{}
	hCamera := &comp.CameraComp{Zoom: 1.0, Speed: 500.0}

	hero := &Hero{
		Entity: ecs.EntityFactory.NewEntity(
			"Player", hRender, hPosition,
			hVelocity, hKeyDirection, hCamera,
		),
	}

	eRender := &comp.RenderComp{
		Sprite: util.GetSprite("res/entity/enemy/monsters.png", 0, 24*8, 16, 24*9),
		Batch:  util.GetBatch("res/entity/enemy/monsters.png"),
		Scale:  3.0,
	}
	ePosition := &comp.PositionComp{X: 128.0, Y: 128.0}

	enemy := &Enemy{
		Entity: ecs.EntityFactory.NewEntity(
			"Player", ePosition, eRender,
		),
	}

	engine.AddSystem(system.NewCamSystem(win))
	engine.AddSystem(system.NewKeySystem(win))
	engine.AddSystem(system.NewMovementSystem())
	engine.AddSystem(system.NewDrawSystem())

	engine.AddEntity(hero)
	engine.AddEntity(enemy)

	engine.Update()
}
