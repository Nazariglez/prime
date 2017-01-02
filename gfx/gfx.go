// Created by nazarigonzalez on 29/12/16.

package gfx

import (
	"errors"
	"log"

	"prime/gfx/gl"
)

var (
	gfxContext *gl.Context

	gfxWidth  int
	gfxHeight int
	gfxTitle  string

	OnStart = func() { log.Println("GFX Event: Start") }
	OnEnd   = func() { log.Println("GFX Event: End") }
	OnDraw  = func() { log.Println("GFX Event: Draw") }
)

func Init(width, height int, title string) error {
	gfxWidth = width
	gfxHeight = height
	gfxTitle = title
	return initialize()
}

func GetContext() (*gl.Context, error) {
	if gfxContext != nil {
		return gfxContext, nil
	}

	return nil, errors.New("Gfx needs start before get the context.")
}
