package imageproc

import (
	"github.com/HeberDamianAlturria/converter-to-braille/internal/braille"
	"image"
	"image/color"
	"image/draw"
	"strings"
)

type ImageProc struct {
	Img image.Image
}

func New(img image.Image) *ImageProc {
	return &ImageProc{Img: img}
}

func (imgproc *ImageProc) applyFloydSteinbergDithering() {
	paletted := image.NewPaletted(imgproc.Img.Bounds(), []color.Color{color.Black, color.White, color.Transparent})

	draw.FloydSteinberg.Draw(paletted, paletted.Bounds(), imgproc.Img, imgproc.Img.Bounds().Min)

	imgproc.Img = paletted
}

func (imgproc *ImageProc) WriteToBrille(strBuilder *strings.Builder, inverted bool) {
	imgproc.applyFloydSteinbergDithering()

	imgBounds := imgproc.Img.Bounds()

	for pixelY := imgBounds.Min.Y; pixelY < imgBounds.Max.Y; pixelY += 4 {
		for pixelX := imgBounds.Min.X; pixelX < imgBounds.Max.X; pixelX += 2 {
			var brailleMatrix braille.BrailleMatrix

			for dotY := 0; dotY < 4; dotY++ {
				for dotX := 0; dotX < 2; dotX++ {
					if pixelY+dotY >= imgBounds.Max.Y || pixelX+dotX >= imgBounds.Max.X {
						continue
					}

					colorAtPixel := imgproc.Img.At(pixelX+dotX, pixelY+dotY)

					brailleMatrix.SetFromColor(dotX, dotY, colorAtPixel, inverted)
				}
			}

			strBuilder.WriteString(brailleMatrix.ToString())
		}
		strBuilder.WriteString("\n")
	}
}
