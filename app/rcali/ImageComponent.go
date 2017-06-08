package rcali

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/jpeg"
	"math"
)

func JpegImage2Bytes(img image.Image) []byte {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	result := buf.Bytes()
	return result
}

func ResizeImage(width, height uint, img image.Image) image.Image {
	dst := resize.Resize(width, height, img, resize.Lanczos3)
	return dst
}

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	//var dx, dy float64 = c.X - x, c.Y - y
	//d := math.Sqrt(dx*dx+dy*dy) / c.R
	//if d > 1 {
	//	return 0
	//} else {
	//	return 255
	//}
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		// outside
		return 0
	} else {
		// inside
		return uint8((1 - math.Pow(d, 5)) * 255)
	}
}

func EmptyIamge(width, height int) image.Image {
	var w, h int = width, height
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	r := 40.0
	θ := 2 * math.Pi / 3
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	cg := &Circle{hw - r*math.Sin(θ), hh - r*math.Cos(θ), 60}
	cb := &Circle{hw - r*math.Sin(-θ), hh - r*math.Cos(-θ), 60}

	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				cr.Brightness(float64(x), float64(y)),
				cg.Brightness(float64(x), float64(y)),
				cb.Brightness(float64(x), float64(y)),
				255,
			}
			m.Set(x, y, c)
		}
	}
	return m
}
