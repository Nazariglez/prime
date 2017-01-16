// Created by nazarigonzalez on 16/1/17.

// +build js

package gfx

import (
	"github.com/gopherjs/gopherjs/js"
)

type Image struct {
	Data   *js.Object
	Width  int
	Height int
}

func NewImage(img *js.Object) *Image {
	return &Image{
		img,
		img.Get("width").Int(),
		img.Get("height").Int(),
	}
}
