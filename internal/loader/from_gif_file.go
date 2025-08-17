package loader

import (
	"image"
	"image/draw"
	"image/gif"
	"os"
)

type ReconstructedFrame struct {
	Image image.Image
	Delay int
}

// FromGifFileReconstructed loads a GIF file and reconstructs its frames into a slice of ReconstructedFrame.
// We need to know that not all the frames of a GIF are a complete image, so we need to manage this scenario.
func FromGifFileReconstructed(path string) ([]ReconstructedFrame, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	g, err := gif.DecodeAll(f)
	if err != nil {
		return nil, err
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
