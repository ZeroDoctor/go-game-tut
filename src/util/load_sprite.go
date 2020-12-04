package util

import (
	"fmt"
	"image"
	"os"

	"github.com/faiface/pixel"

	_ "image/png"
)

var workDir string

func init() {
	var err error
	workDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

type SpriteManager struct {
	sheetMap map[string]pixel.Picture
	batchMap map[string]*pixel.Batch
}

var SM = SpriteManager{
	sheetMap: make(map[string]pixel.Picture),
	batchMap: make(map[string]*pixel.Batch),
}

func GetAllBatches() map[string]*pixel.Batch {
	return SM.batchMap
}

func GetSpriteSheet(path string) (pixel.Picture, error) {
	pic, ok := SM.sheetMap[path]
	if ok {
		return pic, nil
	}

	file, err := os.Open(workDir + "/" + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	pic = pixel.PictureDataFromImage(img)
	SM.sheetMap[path] = pic
	return pic, nil
}

func GetBatch(path string) *pixel.Batch {
	batch, ok := SM.batchMap[path]
	if ok {
		return batch
	}

	spriteSheet, err := GetSpriteSheet(path)
	if err != nil {
		fmt.Println("Failed to load sprite sheet", path, ":\n\t", err)
		return nil
	}
	batch = pixel.NewBatch(&pixel.TrianglesData{}, spriteSheet)
	SM.batchMap[path] = batch
	return batch
}

func GetSprite(path string, x, y, w, h float64) *pixel.Sprite {

	bounds := pixel.Rect{Min: pixel.Vec{X: x, Y: y}, Max: pixel.Vec{X: w, Y: h}}

	spriteSheet, err := GetSpriteSheet(path)
	if err != nil {
		fmt.Println("Failed to load sprite sheet", path, ":\n\t", err)
		return nil
	}
	return pixel.NewSprite(spriteSheet, bounds)
}
