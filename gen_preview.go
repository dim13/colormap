// +build ignore

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/dim13/colormap"
)

func genPreview(name string, p color.Palette, h int) error {
	pimg := genPalette(p)
	img := image.NewRGBA(image.Rect(0, 0, len(p), h))
	for i := 0; i < h; i++ {
		draw.Draw(img, img.Bounds().Add(image.Pt(0, i)), pimg, image.ZP, draw.Src)
	}
	fd, err := os.Create("images/" + name + ".png")
	if err != nil {
		return err
	}
	defer fd.Close()
	return png.Encode(fd, img)
}

func genPalette(p color.Palette) image.Image {
	r := image.Rect(0, 0, len(p), 1)
	img := image.NewPaletted(r, p)
	for i := 0; i < len(p); i++ {
		img.SetColorIndex(i, 0, uint8(i))
	}
	return img
}

func main() {
	genPreview("viridis", colormap.Viridis, 32)
	genPreview("magma", colormap.Magma, 32)
	genPreview("inferno", colormap.Inferno, 32)
	genPreview("plasma", colormap.Plasma, 32)
}
