// Created by nazarigonzalez on 29/12/16.

package gfx

import (
  _"runtime"
  "errors"

  "prime/gfx/gl"
  "log"
)

var (
  gfxContext *gl.Context

  gfxWidth int
  gfxHeight int
  gfxTitle string
  gfxBg []float32

  OnStart = func(){ log.Println("GFX Event: Start") }
  OnEnd = func(){ log.Println("GFX Event: End") }
  OnDraw = func(){ log.Println("GFX Event: Draw") }
)

type GfxEvent int
const (
  GFX_START GfxEvent = iota
  GFX_DRAW
  GFX_END
)

var gfxChannel chan GfxEvent

func Init(width, height int, title string, bg []float32) error {
  gfxWidth = width
  gfxHeight = height
  gfxTitle = title
  gfxBg = bg

  gfxChannel = make(chan GfxEvent)
  //go readEvents()
  return initialize()
}

func GetContext() (*gl.Context, error) {
  if gfxContext != nil {
    return gfxContext, nil
  }

  return nil, errors.New("Gfx needs start before get the context.")
}

func readEvents() {
  for e := range gfxChannel {
    switch e {
    case GFX_START:
      OnStart()
    case GFX_DRAW:
      OnDraw()
    case GFX_END:
      OnEnd()
    }
  }
}
