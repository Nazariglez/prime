// Created by nazarigonzalez on 16/1/17.

// +build !js

package assets

import (
	"image"
	"image/draw"
	_ "image/png"

	"github.com/nazariglez/prime/gfx"

	mAsset "golang.org/x/mobile/asset"
)

func LoadImage(img string) (*gfx.Image, error) {
	f, err := mAsset.Open(img) //os.Open(img)
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
