// Created by nazarigonzalez on 7/1/17.

package gfx

import (
  "prime/gfx/gl"
  "image"
  "image/draw"
)

type Texture struct {
  gl.Texture
}

//https://github.com/cstegel/opengl-samples-golang/blob/master/basic-textures/gfx/texture.go

func NewTexture(img image.Image) {
  rect := img.Bounds()
  rgba := image.NewNRGBA(rect)
  draw.Draw(rgba, rect, img, image.Pt(0,0), draw.Src)
  //todo: stride?

}
