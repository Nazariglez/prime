// Created by nazarigonzalez on 15/1/17.

package gfx

import "image"

type Image struct {
  Data *image.NRGBA
  Width int
  Height int
}

func NewImage(img *image.NRGBA) *Image {
  return &Image{img, img.Rect.Max.X, img.Rect.Max.Y}
}


