package ecs

import (
	"fmt"

	"github.com/ocdogan/rbt"
)

// TODO: use a non-recusive way to traverse trees
// TODO: redo

type ISystem interface {
	Name() string

	Init()
	Update(dt float64)

	AddEntity(entity IEntity)
	RemoveEntity(entity IEntity)
	Requirements() []string
}

type System struct {
	requirements []string
	treeMap      map[string]*rbt.RbTree
}

func NewSystem(requirements ...string) *System {
	return &System{
		requirements: requirements,
		treeMap:      make(map[string]*rbt.RbTree),
	}
}

func (s *System) InspectEntity(e interface{}) IEntity {
	return e.(IEntity)
}

func (s *System) Entities(callback rbt.RbIterationCallback) {
	for _, v := range s.treeMap {
		it, err := v.NewRbIterator(callback)
		if err != nil {
			fmt.Println("Failed to create iterator\n\t", err)
			continue
		}
		it.All()
	}
}

func (s *System) Init() {

}

func (s *System) Update(dt float64) {
	fmt.Println("system update not set")
}

func (s *System) AddEntity(entity IEntity) {
	tree, ok := s.treeMap[entity.GetType()]
	if !ok {
		tree = rbt.NewRbTree()
		s.treeMap[entity.GetType()] = tree
	}

	key := entity.GetID()
	tree.Insert(&key, entity)
}

func (s *System) RemoveEntity(entity IEntity) {
	tree, ok := s.treeMap[entity.GetType()]
	if !ok {
		s.treeMap[entity.GetType()] = rbt.NewRbTree()
	}

	key := entity.GetID()
	tree.Delete(&key)
}

func (s *System) SetRequirements(require string) {
	s.requirements = append(s.requirements, require)
}

func (s *System) Requirements() []string {
	return s.requirements
}
