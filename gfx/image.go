// Created by nazarigonzalez on 15/1/17.

// +build !js

package gfx

import "image"

type Image struct {
	Data   *image.NRGBA
	Width  int
	Height int
}

func NewImage(img *image.NRGBA) *Image {
	return &Image{img, img.Rect.Max.X, img.Rect.Max.Y}
}
