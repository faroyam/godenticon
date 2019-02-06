package godenticon

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
)

// Dermines blocky image size and output image multiplication coefficient
const blockSize = 7
const multiplier = 6

const bgOffset = 10

// GenarateImage creates image based on [16]byte hash and color slice
func GenarateImage(hash [16]byte, colorPalette []color.Color) (*image.RGBA, error) {

	listLength := len(colorPalette)
	if listLength == 0 {
		return nil, errors.New("colors list must not be empty")
	}
	colorNumber := int(hash[0])
	colorFromhash := colorPalette[(colorNumber % listLength)]
	if colorFromhash == nil {
		return nil, errors.New("color must not be nil")
	}

	blocksFromhash := [blockSize][blockSize]bool{
		{hash[0]&1 != 0, hash[0]&128 != 0, hash[1]&64 != 0, hash[2]&32 != 0, hash[1]&64 != 0, hash[0]&128 != 0, hash[0]&1 != 0},
		{hash[0]&2 != 0, hash[1]&1 != 0, hash[1]&128 != 0, hash[2]&64 != 0, hash[1]&128 != 0, hash[1]&1 != 0, hash[0]&2 != 0},
		{hash[0]&4 != 0, hash[1]&2 != 0, hash[2]&1 != 0, hash[2]&128 != 0, hash[2]&1 != 0, hash[1]&2 != 0, hash[0]&4 != 0},
		{hash[0]&8 != 0, hash[1]&4 != 0, hash[2]&2 != 0, hash[3]&1 != 0, hash[2]&2 != 0, hash[1]&4 != 0, hash[0]&8 != 0},
		{hash[0]&16 != 0, hash[1]&8 != 0, hash[2]&4 != 0, hash[3]&2 != 0, hash[2]&4 != 0, hash[1]&8 != 0, hash[0]&16 != 0},
		{hash[0]&32 != 0, hash[1]&16 != 0, hash[2]&8 != 0, hash[3]&32 != 0, hash[2]&8 != 0, hash[1]&16 != 0, hash[0]&32 != 0},
		{hash[0]&64 != 0, hash[1]&32 != 0, hash[2]&16 != 0, hash[3]&64 != 0, hash[2]&16 != 0, hash[1]&32 != 0, hash[0]&64 != 0},
	}

	imageSize := blockSize * multiplier

	img := func() *image.RGBA {
		maxPoint := image.Point{imageSize, imageSize}
		return image.NewRGBA(image.Rectangle{image.ZP, maxPoint})
	}()

	for x := 0; x < imageSize; x++ {
		for y := 0; y < imageSize; y++ {
			if blocksFromhash[int(y/multiplier)][int(x/multiplier)] {
				img.Set(x, y, colorFromhash)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	bg := func() *image.RGBA {
		maxPoint := image.Point{imageSize + bgOffset, imageSize + bgOffset}
		return image.NewRGBA(image.Rectangle{image.ZP, maxPoint})
	}()

	offset := image.Pt(bgOffset/2, bgOffset/2)

	draw.Draw(bg, bg.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	draw.Draw(bg, img.Bounds().Add(offset), img, image.ZP, draw.Src)

	return bg, nil
}

// Sample color slices collected in palettes

// ColorBlues ...
var ColorBlues = []color.Color{
	color.RGBA{0, 0, 32, 0xff},
	color.RGBA{0, 32, 64, 0xff},
	color.RGBA{0, 64, 96, 0xff},
	color.RGBA{0, 96, 128, 0xff},
	color.RGBA{0, 112, 144, 0xff},
	color.RGBA{63, 104, 150, 0xff},
	color.RGBA{106, 151, 193, 0xff},
	color.RGBA{138, 173, 221, 0xff},
	color.RGBA{32, 67, 116, 0xff},
	color.RGBA{143, 185, 255, 0xff},
}

// ColorNeon ...
var ColorNeon = []color.Color{
	color.RGBA{2, 254, 0, 0xff},
	color.RGBA{0, 224, 0, 0xff},
	color.RGBA{0, 192, 0, 0xff},
	color.RGBA{0, 192, 0, 0xff},
	color.RGBA{0, 160, 0, 0xff},
	color.RGBA{0, 128, 0, 0xff},
	color.RGBA{126, 255, 0, 0xff},
	color.RGBA{70, 195, 76, 0xff},
	color.RGBA{71, 159, 120, 0xff},
	color.RGBA{207, 255, 177, 0xff},
}

// ColorDiscord ...
var ColorDiscord = []color.Color{
	color.RGBA{128, 128, 128, 0xff},
	color.RGBA{96, 96, 96, 0xff},
	color.RGBA{64, 64, 64, 0xff},
	color.RGBA{32, 32, 32, 0xff},
	color.RGBA{1, 1, 1, 0xff},
	color.RGBA{72, 75, 81, 0xff},
	color.RGBA{54, 57, 63, 0xff},
	color.RGBA{47, 49, 54, 0xff},
	color.RGBA{42, 44, 49, 0xff},
	color.RGBA{32, 34, 37, 0xff},
}

// ColorSunny ...
var ColorSunny = []color.Color{
	color.RGBA{241, 220, 6, 0xff},
	color.RGBA{244, 244, 10, 0xff},
	color.RGBA{244, 255, 0, 0xff},
	color.RGBA{255, 210, 0, 0xff},
	color.RGBA{233, 243, 15, 0xff},
	color.RGBA{249, 156, 0, 0xff},
	color.RGBA{255, 167, 20, 0xff},
	color.RGBA{255, 202, 22, 0xff},
	color.RGBA{255, 178, 27, 0xff},
	color.RGBA{249, 115, 0, 0xff},
}

// ColorGopher ...
var ColorGopher = []color.Color{
	color.RGBA{100, 200, 200, 0xff},
}
