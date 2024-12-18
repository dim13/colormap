//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"image/color"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type floatColor struct {
	r, g, b float64
}

func readPalette(fname string) ([]floatColor, error) {
	fd, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	var fc []floatColor
	for {
		var f floatColor
		if _, err := fmt.Fscanf(fd, "%f %f %f", &f.r, &f.g, &f.b); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		fc = append(fc, f)
	}
	return fc, nil
}

func (f floatColor) RGBA() (r uint32, g uint32, b uint32, a uint32) {
	var x float64 = 0xffff
	r = uint32(x * f.r)
	g = uint32(x * f.g)
	b = uint32(x * f.b)
	a = uint32(x * 1.0)
	return
}

func gen(fname string) error {
	p, err := readPalette(fname)
	if err != nil {
		return err
	}
	name := strings.TrimSuffix(filepath.Base(fname), filepath.Ext(fname))
	fd, err := os.Create(name + ".go")
	if err != nil {
		return err
	}
	defer fd.Close()
	varName := strings.Title(name)
	fmt.Fprintf(fd, "// Code generated by gen_palette.go DO NOT EDIT.\n")
	fmt.Fprintf(fd, "package colormap\n\n")
	fmt.Fprintf(fd, "import \"image/color\"\n\n")
	fmt.Fprintf(fd, "// %s palette\n", varName)
	fmt.Fprintf(fd, "var %s = color.Palette{\n", varName)
	for _, v := range p {
		c := color.RGBAModel.Convert(v).(color.RGBA)
		fmt.Fprintf(fd, "\tcolor.RGBA{%#02x, %#02x, %#02x, %#02x},\n", c.R, c.G, c.B, c.A)
	}
	fmt.Fprintf(fd, "}\n")
	return nil
}

func main() {
	gen("palette/magma.txt")
	gen("palette/inferno.txt")
	gen("palette/plasma.txt")
	gen("palette/viridis.txt")
}
