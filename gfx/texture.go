// Created by nazarigonzalez on 7/1/17.

package gfx

import (
	"log"

	"github.com/nazariglez/prime/gfx/gl"
)

var countID = 0

type Texture struct {
	ID 		int
	Tex    *gl.Texture
	Width  float32
	Height float32
}

func NewTexture(img *Image) *Texture {
	if img == nil {
		log.Fatal("Error creating a new texture, img cannot be nil.")
	}

	tex := GL.CreateTexture()
	GL.BindTexture(GL.TEXTURE_2D, tex)

	GL.TexParameteri(GL.TEXTURE_2D, GL.TEXTURE_WRAP_S, GL.CLAMP_TO_EDGE)
	GL.TexParameteri(GL.TEXTURE_2D, GL.TEXTURE_WRAP_T, GL.CLAMP_TO_EDGE)
	GL.TexParameteri(GL.TEXTURE_2D, GL.TEXTURE_MIN_FILTER, GL.LINEAR) //todo change to nearest in options
	GL.TexParameteri(GL.TEXTURE_2D, GL.TEXTURE_MAG_FILTER, GL.LINEAR)

	GL.TexImage2D(GL.TEXTURE_2D, 0, GL.RGBA, GL.RGBA, GL.UNSIGNED_BYTE, img.Data)

	countID++
	return &Texture{countID, tex, float32(img.Width), float32(img.Height)}
}
