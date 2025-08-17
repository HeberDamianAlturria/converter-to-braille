package imageproc

import (
	"github.com/HeberDamianAlturria/converter-to-braille/internal/braille"
	"image"
	"image/color"
	"image/draw"
	"strings"
)

type ImageProc struct {
	img image.Image
}

func New(img image.Image) *ImageProc {
	return &ImageProc{img: img}
}

func (imgproc *ImageProc) applyFloydSteinbergDithering() {
	paletted := image.NewPaletted(imgproc.img.Bounds(), []color.Color{color.Black, color.White, color.Transparent})

	draw.FloydSteinberg.Draw(paletted, paletted.Bounds(), imgproc.img, imgproc.img.Bounds().Min)

	imgproc.img = paletted
}

func (imgproc *ImageProc) WriteToBrille(strBuilder *strings.Builder) {
	imgproc.applyFloydSteinbergDithering()

	imgBounds := imgproc.img.Bounds()

	for pixelY := imgBounds.Min.Y; pixelY < imgBounds.Max.Y; pixelY += 4 {
		for pixelX := imgBounds.Min.X; pixelX < imgBounds.Max.X; pixelX += 2 {
			var brailleMatrix braille.BrailleMatrix

			for dotY := 0; dotY < 4; dotY++ {
				for dotX := 0; dotX < 2; dotX++ {
					if pixelY+dotY >= imgBounds.Max.Y || pixelX+dotX >= imgBounds.Max.X {
						continue
					}

					colorAtPixel := imgproc.img.At(pixelX+dotX, pixelY+dotY)

					if colorAtPixel == color.Black {
						brailleMatrix[dotX][dotY] = 0
					} else {
						brailleMatrix[dotX][dotY] = 1
					}
				}
			}

			strBuilder.WriteString(brailleMatrix.ToString())
		}
		strBuilder.WriteString("\n")
	}
}
