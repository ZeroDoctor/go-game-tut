package main

import (
	"os"

	"github.com/faiface/pixel/pixelgl"
	"github.com/zerodoctor/go-tut/game"
	test "github.com/zerodoctor/go-tut/main-test"
)

func main() {

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "-test" {
		test.Route(args[1])
		return
	}

	pixelgl.Run(game.Run)
}
