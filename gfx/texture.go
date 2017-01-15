// Created by nazarigonzalez on 7/1/17.

package gfx

import (
  "prime/gfx/gl"
)

type Texture struct {
  Tex *gl.Texture
  Width float32
  Height float32
}

func NewTexture(img *Image) *Texture {
  tex := GL.CreateTexture()
  GL.BindTexture(GL.TEXTURE_2D, tex)

  GL.TexParameteri(GL.TEXTURE_2D, GL.TEXTURE_WRAP_S, GL.CLAMP_TO_EDGE)
  GL.TexParameteri(GL.TEXTURE_2D, GL.TEXTURE_WRAP_T, GL.CLAMP_TO_EDGE)
  GL.TexParameteri(GL.TEXTURE_2D, GL.TEXTURE_MIN_FILTER, GL.LINEAR)
  GL.TexParameteri(GL.TEXTURE_2D, GL.TEXTURE_MAG_FILTER, GL.NEAREST)

  //todo: manage nil img.Data?

  GL.TexImage2D(GL.TEXTURE_2D, 0, GL.RGBA, GL.RGBA, GL.UNSIGNED_BYTE, img.Data)

  return &Texture{tex, float32(img.Width), float32(img.Height)}
}
