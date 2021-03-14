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
	sheetMap map[string]*SpriteSheet
	batchMap map[string]*pixel.Batch
}

var SM = SpriteManager{
	sheetMap: make(map[string]*SpriteSheet),
	batchMap: make(map[string]*pixel.Batch),
}

func GetAllBatches() map[string]*pixel.Batch {
	return SM.batchMap
}

type SpriteSheet struct {
	pic        *pixel.PictureData
	tileWidth  float64
	tileHeight float64
	width      float64
	height     float64
}

func LoadSpriteSheet(path, name string, tileWidth, tileHeight float64) error {
	_, ok := SM.sheetMap[name]
	if ok {
		fmt.Println("WARN: spritesheet", name, "is already loaded")
		return nil
	}

	file, err := os.Open(workDir + "/" + path)
	if err != nil {
		return fmt.Errorf("ERROR: in opening path %s:\n\t%s", path, err.Error())
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("ERROR: in decoding path %s:\n\t%s", path, err.Error())
	}
	pic := pixel.PictureDataFromImage(img)

	sheet := &SpriteSheet{
		pic:        pic,
		tileWidth:  tileWidth,
		tileHeight: tileHeight,
		width:      pic.Bounds().W(),
		height:     pic.Bounds().H(),
	}

	SM.sheetMap[name] = sheet

	return nil
}

func getSpriteSheet(name string) (*SpriteSheet, error) {
	sheet, ok := SM.sheetMap[name]
	if !ok {
		return nil, fmt.Errorf("ERROR: failed to find sprite sheet with name: %s", name)
	}

	return sheet, nil
}

func GetBatch(name string) (*pixel.Batch, error) {
	batch, ok := SM.batchMap[name]
	if ok {
		return batch, nil
	}

	spriteSheet, err := getSpriteSheet(name)
	if err != nil {
		return nil, err
	}

	batch = pixel.NewBatch(&pixel.TrianglesData{}, spriteSheet.pic)
	SM.batchMap[name] = batch
	return batch, nil
}

func GetSprite(name string, x1, y1, x2, y2 float64) *pixel.Sprite {
	spriteSheet, err := getSpriteSheet(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	trueHeight := spriteSheet.height / spriteSheet.tileHeight

	bounds := pixel.Rect{
		Min: pixel.Vec{X: spriteSheet.tileWidth * x1, Y: spriteSheet.tileHeight * (trueHeight - y2)},
		Max: pixel.Vec{X: spriteSheet.tileWidth * x2, Y: spriteSheet.tileHeight * (trueHeight - y1)},
	}

	fmt.Println(spriteSheet.pic.Bounds(), bounds)
	return pixel.NewSprite(spriteSheet.pic, bounds)
}
