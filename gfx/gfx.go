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

var lockChannel chan func()

var (
	GLContext *gl.Context

	gfxWidth  int
	gfxHeight int
	gfxScale 	int
	gfxTitle  string

	OnStart = func() { log.Println("GFX Event: Start") }
	OnEnd   = func() { log.Println("GFX Event: End") }
)

func Init(width, height int, title string, scale int) error {
	gfxWidth = width
	gfxHeight = height
	gfxTitle = title
	gfxScale = scale

	lockChannel = make(chan func())

	return initialize()
}

func Render(f func() error) { //todo pass scene graph
	err := RunSafeFn(func() error {
		if err := f(); err != nil {
			return err
		}

		postRender()
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func RunSafeFn(f func() error) error {
	if lockChannel == nil {
		return errors.New("Invalid Safe function. (Channel closed)")
	}

	//reminder: sync.Wait has some problem with runtime.LockOSThread
	w := make(chan bool) //todo sync.Pool this
	var err error

	lockChannel <- func() {
		err = f()
		w <- true
	}

	<-w
	close(w)

	return err
}