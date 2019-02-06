package main

import (
	"crypto/md5"
	"image/png"
	"os"

	g "github.com/faroyam/godenticon"
)

func main() {
	// Calculate MD5 checksum of the data
	hash := md5.Sum([]byte("Gopher!"))

	// Create image using hash and slice of colors
	img, _ := g.GenarateImage(hash, g.ColorGopher)

	// Create new file to be written
	f, _ := os.Create("image.png")

	// Write image to file
	png.Encode(f, img)
}
