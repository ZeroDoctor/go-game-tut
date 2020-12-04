package ecs

import (
	"fmt"
	"math/rand"

	"github.com/ocdogan/rbt"
)

const (
	ELeft EDir = iota
	ERight
)

type EDir int

type EntityManager struct {
	nextid rbt.Uint32Key
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		nextid: 0,
	}
}

func (e *EntityManager) NewEntity(eType string, comps ...IComponent) *Entity {
	e.nextid++
	e.nextid = rbt.Uint32Key(uint32(e.nextid) + uint32(rand.Intn(6)))

	compName := ""
	for _, c := range comps {
		compName += (c.Name() + " ")
	}

	return &Entity{
		id:         e.nextid,
		eType:      eType,
		components: comps,
		compName:   compName[:len(compName)-1],
	}
}

var EntityFactory = NewEntityManager()

type IEntity interface {
	GetID() rbt.Uint32Key
	AddComponent(component IComponent)
	GetComponent(name string) IComponent
	GetAllComponents() string
	GetChildren() []IEntity
	GetChild(direction EDir) IEntity
	AddChild(entity IEntity, direction EDir)
	GetType() string
}

type Entity struct {
	id         rbt.Uint32Key
	eType      string
	components []IComponent
	children   []IEntity
	compName   string
}

func (e Entity) GetID() rbt.Uint32Key {
	return e.id
}

func (e *Entity) AddComponent(component IComponent) {
	e.components = append(e.components, component)

	name := component.Name()
	if e.compName != "" {
		name = " " + component.Name()
	}

	e.compName += name
}

func (e *Entity) GetComponent(name string) IComponent {
	for _, c := range e.components {
		if c.Name() == name {
			return c
		}
	}

	fmt.Println("ERROR: component", name, "not found")
	return nil
}

func (e *Entity) GetAllComponents() string {
	return e.compName
}

func (e *Entity) GetChildren() []IEntity {
	return e.children
}

func (e *Entity) GetChild(direction EDir) IEntity {
	return e.children[direction]
}

func (e *Entity) AddChild(entity IEntity, dir EDir) {
	e.children[dir] = entity
}

func (e *Entity) GetType() string {
	return e.eType
}
