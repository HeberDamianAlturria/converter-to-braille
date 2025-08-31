package loader

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strings"
)

func GetImage(path string) (image.Image, error) {
	var (
		img      image.Image
		imgError error
	)

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		response, err := http.Get(path)

		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		img, _, imgError = image.Decode(response.Body)
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		img, _, imgError = image.Decode(file)
	}

	if imgError != nil {
		return nil, imgError
	}

	return img, nil
}
