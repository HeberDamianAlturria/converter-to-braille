package loader

import (
	"image/gif"
	"os"
)

func FromGifFile(path string) (*gif.GIF, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	g, err := gif.DecodeAll(f)

	if err != nil {
		return nil, err
	}

	return g, nil

}
