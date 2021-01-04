package system

import (
	"github.com/ocdogan/rbt"
	"github.com/zerodoctor/go-tut/ecs"
)

type GenerateSystem struct {
	name string
	*ecs.System
}

func NewGenerateSystem() *GenerateSystem {
	return &GenerateSystem{
		name:   "GenerateSystem",
		System: ecs.NewSystem("render"),
	}
}

func (m *GenerateSystem) Init() {

}

func (m *GenerateSystem) Update(dt float64) {
	m.Entities(func(iterator rbt.RbIterator, key rbt.RbKey, value interface{}) {

	})
}
