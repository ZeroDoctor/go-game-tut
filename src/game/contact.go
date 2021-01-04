package game

import (
	"fmt"

	"github.com/ByteArena/box2d"
	"github.com/zerodoctor/go-tut/ecs"
)

type ContactListener struct {
}

func (c *ContactListener) BeginContact(contact box2d.B2ContactInterface) { // contact has to be backed by a pointer
	entityA, okA := contact.GetFixtureA().GetUserData().(ecs.IEntity)
	entityB, okB := contact.GetFixtureB().GetUserData().(ecs.IEntity)
	if okA && okB {
		fmt.Println(entityA.GetType(), entityB.GetType())

	}

}

func (c *ContactListener) EndContact(contact box2d.B2ContactInterface) { // contact has to be backed by a pointer

}

func (c *ContactListener) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) { // contact has to be backed by a pointer

}

func (c *ContactListener) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) { // contact has to be backed by a pointer

}
