// Created by nazarigonzalez on 29/12/16.

// +build !js
// +build !android

package gfx

import (
  "github.com/go-gl/glfw/v3.2/glfw"
  "log"
  "runtime"

  "prime/gfx/gl"
)

func initialize() error {
  runtime.LockOSThread()
  log.Println("Desktop initialized")

  if err := glfw.Init(); err != nil {
    return err
  }

  defer glfw.Terminate()

  glfw.WindowHint(glfw.Samples, 4)
  glfw.WindowHint(glfw.ContextVersionMajor, 2)
  glfw.WindowHint(glfw.ContextVersionMinor, 1)

  window, err := glfw.CreateWindow(gfxWidth, gfxHeight, gfxTitle, nil, nil)
  if err != nil {
    return err
  }

  window.MakeContextCurrent()

  ctx, err := gl.NewContext()
  if err != nil {
    return err
  }
  gfxContext = ctx

  OnStart() //todo: pass the ctx as argument?

  for !window.ShouldClose() {
    //draw here
    OnDraw()

    window.SwapBuffers()
    glfw.PollEvents()
  }

  OnEnd()
  return nil
}
