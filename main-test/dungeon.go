package test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Tree struct {
	x, y          float64
	width, height float64
	room          bool
}

func generation() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	const div = 32
	var tree [div]Tree
	tree[0] = Tree{x: 0, y: 0, width: 1024 - (0 * 2), height: 768 - (0 * 2)}
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 1; i < (div - 1); i += 2 {
		parentIndex := int(i / 2)
		parent := tree[parentIndex]

		left := Tree{}
		right := Tree{}

		now := time.Now()
		randVert := rand.Intn(int(now.Unix()%2000)+1000) * (rand.Intn(6) + 1)
		// TODO: weight random generation
		//	create ratio, and size limit var

		fmt.Println(parent)
		width := float64(rand.Intn(int(parent.width/3))) + parent.width/3
		height := float64(rand.Intn(int(parent.height/4))) + parent.height/4
		fmt.Println(width, height)
		vert := randVert%2 == 0
		if vert && width < 100 {
			vert = !vert
		} else if !vert && height < 100 {
			vert = !vert
		}

		if vert {
			left.x = parent.x
			left.y = parent.y
			left.width = width
			left.height = parent.height

			right.x = parent.x + left.width
			right.y = parent.y
			right.width = parent.width - left.width
			right.height = parent.height
		} else {
			left.x = parent.x
			left.y = parent.y
			left.width = parent.width
			left.height = height

			right.x = parent.x
			right.y = parent.y + height
			right.width = parent.width
			right.height = parent.height - height
		}

		fmt.Println("\t", left, right)

		tree[i] = left
		tree[i+1] = right
	}

	rooms := []Tree{}

	for i := len(tree) - 1; i >= 0; i-- {
		r := Tree{}

		// TODO: random width and height with limit
		// 	random x, and y with the limit of half the leaf x, and y
		//  find diection of rooms and set a predetermine width and height
		// 	depending on direction and set the other width or height to the center
		// 	of the destination leaf

		rooms = append(rooms, r)
	}

	var vectors []pixel.Vec
	imd := imdraw.New(nil)
	imd.Color = colornames.Black
	imd.EndShape = imdraw.RoundEndShape

	for _, t := range tree {
		vec := CreateVertex(t.x, t.y, t.width, t.height)
		temp := []pixel.Vec{pixel.V(vec[0].X, vec[0].Y), pixel.V(vec[1].X, vec[1].Y)}
		vectors = append(vectors, temp...)
	}

	fmt.Println(vectors)
	current := 0
	count := 0

	for !win.Closed() {

		if win.JustPressed(pixelgl.KeyN) && len(vectors) > current+2 {
			current += 2
			count++
			fmt.Println("Clicked")
			fmt.Println(vectors[current-2 : current])
			imd.Color = colornames.Black
			if count%2 == 0 {
				imd.Color = colornames.Blue
			}
			imd.Push(vectors[current-2 : current]...)
			imd.Rectangle(2.0)
		}

		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		win.Update()
		if win.JustPressed(pixelgl.KeyEscape) {
			win.Destroy()
			return
		}
	}
}

func CreateVertex(x, y, w, h float64) []pixel.Vec {
	return []pixel.Vec{
		{X: x, Y: y},
		{X: x + w, Y: y + h},
	}
}

func TestDungeon() {
	pixelgl.Run(generation)
}
