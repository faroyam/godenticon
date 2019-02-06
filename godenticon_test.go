package godenticon_test

import (
	"crypto/md5"
	"testing"

	g "github.com/faroyam/godenticon"
)

func TestNilList(t *testing.T) {
	if _, err := g.GenarateImage(md5.Sum([]byte("Golang!")), nil); err == nil {
		t.Errorf("no error with nil color list")
	}

}

func TestNilHash(t *testing.T) {
	if _, err := g.GenarateImage([16]byte{}, g.ColorGopher); err != nil {
		t.Errorf(err.Error())
	}
}

func TestSymmetry(t *testing.T) {
	img, err := g.GenarateImage(md5.Sum([]byte("Golang!")), g.ColorGopher)
	if err != nil {
		t.Errorf(err.Error())
	}

	for x := 0; x < img.Rect.Max.X; x++ {
		for y := 0; y < img.Rect.Max.Y; y++ {
			if img.At(x, y) != img.At(img.Rect.Max.X-1-x, y) {
				t.Errorf("image lacks vertical symmetry")
			}
		}
	}
}
