// Created by nazarigonzalez on 16/1/17.

// +build !js

package assets

import (
	"image"
	"image/draw"
	_ "image/png"
	"os"

	"prime/gfx"
)

func LoadImage(img string) (*gfx.Image, error) {
	f, err := os.Open(img)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	i, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	bounds := i.Bounds()
	rgba := image.NewNRGBA(bounds)
	draw.Draw(rgba, rgba.Bounds(), i, bounds.Min, draw.Src)

	return gfx.NewImage(rgba), nil
}
