// Created by nazarigonzalez on 29/12/16.

// +build !js
// +build !android

package gfx

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"

	"prime/gfx/gl"
)

var win *glfw.Window

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

	win = window

	window.MakeContextCurrent()

	ctx, err := gl.NewContext()
	if err != nil {
		return err
	}

	GLContext = ctx

	go OnStart()

	return RunSafeReader()
}

func postRender() {
	if win.ShouldClose() {
		Close(nil)
	}

	win.SwapBuffers()
	glfw.PollEvents()
}
