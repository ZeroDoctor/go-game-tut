package ecs

import (
	"fmt"
	"strings"
	"time"

	"github.com/ByteArena/box2d"
	"github.com/zerodoctor/go-tut/src/util"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel/pixelgl"
)

type Engine struct {
	Win      *pixelgl.Window
	Cfg      pixelgl.WindowConfig
	World    *box2d.B2World
	systems  []ISystem
	entities []IEntity
}

func NewEngine(Win *pixelgl.Window, Cfg pixelgl.WindowConfig, World *box2d.B2World) *Engine {
	return &Engine{
		Win:   Win,
		Cfg:   Cfg,
		World: World,
	}
}

func (e *Engine) AddEntity(entity IEntity) {
	e.entities = append(e.entities, entity)

outter:
	for _, s := range e.systems {
		shouldAdd := false
		for _, req := range s.Requirements() {
			shouldAdd = strings.Contains(entity.GetAllComponents(), req)
			if !shouldAdd {
				continue outter
			}
		}
		s.AddEntity(entity)
	}
}

func (e *Engine) AddEntityToSystem(entity IEntity, system string) {

}

func (e *Engine) AddComponent(entity IEntity, component IComponent) {
	for _, e := range e.entities {
		if e.GetID() == entity.GetID() {
			e.AddComponent(component)
		}
	}
}

func (e *Engine) RemoveEntity(entity IEntity) {
	for _, s := range e.systems {
		s.RemoveEntity(entity)
	}
}

func (e *Engine) AddSystem(system ISystem) {
	e.systems = append(e.systems, system)
}

func (e *Engine) Update() {
	frames := 0
	second := time.Tick(1 * time.Second)

	for _, s := range e.systems {
		s.Init()
	}

	last := time.Now()
	for !e.Win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		e.Win.Clear(colornames.Slateblue)

		for _, s := range e.systems {
			s.Update(dt)
		}

		e.World.Step(1.0/60.0, 8.0, 3.0)

		for _, b := range util.GetAllBatches() {
			b.Draw(e.Win)
			b.Clear()
		}

		e.Win.Update()

		frames++
		select {
		case <-second:
			e.Win.SetTitle(fmt.Sprintf("%s | FPS: %d", e.Cfg.Title, frames))
			frames = 0
		default:
		}
	}
}
