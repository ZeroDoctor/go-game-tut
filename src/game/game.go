package game

import (
	"github.com/ByteArena/box2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/zerodoctor/go-tut/ecs"
	"github.com/zerodoctor/go-tut/game/comp"
	"github.com/zerodoctor/go-tut/game/system"
	"github.com/zerodoctor/go-tut/util"
)

func Run() {

	err := util.LoadSpriteSheet("res/entity/wizard_player.png", "player.png", 64.0, 64.0)
	if err != nil {
		panic(err)
	}
	err = util.LoadSpriteSheet("res/entity/enemy/monsters.png", "monsters.png", 16.0, 24.0)
	if err != nil {
		panic(err)
	}

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

	world.SetContactListener(&ContactListener{})

	engine := ecs.NewEngine(win, cfg, &world)

	hero := &Hero{}
	hbody := CreateBox2DDefault(
		&world, hero,
		0.0, 0.0, 32.0, 32.0,
		1.0, 0.3,
	)

	hbatch, err := util.GetBatch("player.png")
	if err != nil {
		panic(err)
	}
	hRender := &comp.RenderComp{
		Sprite: util.GetSprite("player.png", 0, 2, 1, 3),
		Batch:  hbatch,
		Scale:  1.5,
	}
	hPosition := &comp.PositionComp{Body: hbody}
	hVelocity := &comp.VelocityComp{Maxv: 0.75, Minv: -0.75, Speed: 0.9}
	hKeyDirection := &comp.DirectionComp{}
	hCamera := &comp.CameraComp{Zoom: 1.0, Speed: 500.0}

	hero.Entity = ecs.EntityFactory.NewEntity(
		"Player", hRender, hPosition,
		hVelocity, hKeyDirection, hCamera,
	)

	enemy := &Enemy{}
	ebody := CreateBox2DDefault(
		&world, enemy,
		128.0, 128.0, 16.0, 24.0,
		1.0, 0.3,
	)

	ebatch, err := util.GetBatch("monsters.png")
	if err != nil {
		panic(err)
	}
	eRender := &comp.RenderComp{
		Sprite: util.GetSprite("monsters.png", 0, 1, 1, 2),
		Batch:  ebatch,
		Scale:  3.0,
	}
	eVelocity := &comp.VelocityComp{Maxv: 0.95, Minv: -0.95, Speed: 0.9}
	ePosition := &comp.PositionComp{Body: ebody}

	enemy.Entity = ecs.EntityFactory.NewEntity(
		"Enemy", ePosition, eVelocity, eRender,
	)

	enemy1 := &Enemy{}
	ebody1 := CreateBox2DDefault(
		&world, enemy1,
		-128.0, -128.0, 16.0, 24.0,
		1.0, 0.3,
	)

	ebatch1, err := util.GetBatch("monsters.png")
	if err != nil {
		panic(err)
	}
	eRender1 := &comp.RenderComp{
		Sprite: util.GetSprite("monsters.png", 0, 2, 1, 3),
		Batch:  ebatch1,
		Scale:  3.0,
	}
	ePosition1 := &comp.PositionComp{Body: ebody1}
	eVelocity1 := &comp.VelocityComp{Maxv: 0.95, Minv: -0.95, Speed: 0.9}

	enemy1.Entity = ecs.EntityFactory.NewEntity(
		"Enemy1", eVelocity1, ePosition1, eRender1,
	)

	engine.AddSystem(system.NewCamSystem(win))
	engine.AddSystem(system.NewKeySystem(win))
	engine.AddSystem(system.NewMovementSystem())
	engine.AddSystem(system.NewDrawSystem())

	engine.AddEntity(hero)
	engine.AddEntity(enemy)
	engine.AddEntity(enemy1)

	engine.Update()
}

func CreateBox2DDefault(world *box2d.B2World, entity ecs.IEntity, x, y, w, h float64, density, friction float64) *box2d.B2Body {
	xm := util.PixelToMeters(x)
	ym := util.PixelToMeters(y)
	wm := util.PixelToMeters(w)
	hm := util.PixelToMeters(h)

	bodydef := box2d.NewB2BodyDef()
	bodydef.Type = box2d.B2BodyType.B2_dynamicBody
	bodydef.Position.Set(xm, ym)
	body := world.CreateBody(bodydef)

	poly := box2d.B2PolygonShape{}
	poly.SetAsBox(wm, hm)

	fixdef := box2d.B2FixtureDef{}
	fixdef.UserData = entity
	fixdef.Shape = &poly
	fixdef.Density = density
	fixdef.Friction = friction

	body.CreateFixtureFromDef(&fixdef)

	return body
}
