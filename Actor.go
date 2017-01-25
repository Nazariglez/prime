// Created by nazarigonzalez on 15/1/17.

package prime

import "github.com/nazariglez/prime/gfx"

//todo actor pool

type Actor struct {
	texture  *gfx.Texture
	bounds   *Rect
	position *Point
	scale    *Point
}

func (a *Actor) Texture() *gfx.Texture {
	return a.texture
}

func (a *Actor) Bounds() *Rect {
	return a.bounds
}

func (a *Actor) Position() *Point {
	return a.position
}

func (a *Actor) Scale() *Point {
	return a.scale
}

type GameObject interface {
	Texture() *gfx.Texture
	Bounds() *Rect
	Scale() *Point
	Position() *Point
}

type Rect struct {
	X, Y, Width, Height int
}

type Point struct {
	X, Y int
}
