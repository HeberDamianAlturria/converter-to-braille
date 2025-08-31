package loader

import (
	"image"
	"strings"
)

func GetImage(path string) (image.Image, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return FromImageURL(path)
	}
	return FromImageFile(path)
}
