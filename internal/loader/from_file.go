package loader

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func FromFile(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
