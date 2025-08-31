package loader

import (
	"image"
	"image/draw"
	"image/gif"
	"net/http"
	"os"
	"strings"
)

type ReconstructedFrame struct {
	Image image.Image
	Delay int
}

// GetGifReconstructed loads a GIF from a file path or URL and reconstructs its frames into a slice of ReconstructedFrame.
func GetGifReconstructed(path string) ([]ReconstructedFrame, error) {
	var (
		g        *gif.GIF
		gifError error
	)

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {

		response, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		g, gifError = gif.DecodeAll(response.Body)
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		g, gifError = gif.DecodeAll(file)
	}

	if gifError != nil {
		return nil, gifError
	}

	canvas := image.NewPaletted(g.Image[0].Bounds(), g.Image[0].Palette) // Canvas to draw frames on
	frames := make([]ReconstructedFrame, len(g.Image))

	for i, frame := range g.Image {
		if i > 0 && g.Disposal[i-1] == gif.DisposalBackground {
			// Clear the canvas if the previous frame's disposal method is to clear the background.
			draw.Draw(canvas, canvas.Bounds(), &image.Uniform{g.Image[0].Palette[g.BackgroundIndex]}, image.Point{}, draw.Src)
		}

		// Draw the current frame onto the canvas.
		draw.Draw(canvas, frame.Bounds(), frame, image.Point{}, draw.Over)

		// Create a copy of the canvas for the frame.
		copyCanvas := image.NewPaletted(canvas.Bounds(), canvas.Palette)
		draw.Draw(copyCanvas, copyCanvas.Bounds(), canvas, image.Point{}, draw.Src)

		frames[i] = ReconstructedFrame{Image: copyCanvas, Delay: g.Delay[i]}
	}

	return frames, nil
}
