// Created by nazarigonzalez on 29/12/16.

package gfx

import (
	"errors"
	"log"

	"prime/gfx/gl"
)

const (
	BROWSER_SCALE_NONE int = iota+1
	BROWSER_SCALE_FIT
	BROWSER_SCALE_FILL
	BROWSER_SCALE_ASPECT_FILL
)

var (
	gfxContext *gl.Context

	gfxWidth  int
	gfxHeight int
	gfxScale 	int
	gfxTitle  string

	OnStart = func() { log.Println("GFX Event: Start") }
	OnEnd   = func() { log.Println("GFX Event: End") }
	OnDraw  = func() { log.Println("GFX Event: Draw") }
)

func Init(width, height int, title string, scale int) error {
	gfxWidth = width
	gfxHeight = height
	gfxTitle = title
	gfxScale = scale
	return initialize()
}

func GetContext() (*gl.Context, error) {
	if gfxContext != nil {
		return gfxContext, nil
	}

	return nil, errors.New("Gfx needs start before get the context.")
}
