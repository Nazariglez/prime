// Created by nazarigonzalez on 29/12/16.

package gfx

import (
	"errors"
	"log"

	"github.com/nazariglez/prime/gfx/gl"
)

const (
	BROWSER_SCALE_NONE int = iota + 1
	BROWSER_SCALE_FIT
	BROWSER_SCALE_FILL
	BROWSER_SCALE_ASPECT_FILL
)

var lockChannel chan func()
var closeChannel chan error

var (
	GL *gl.Context

	gfxWidth  int
	gfxHeight int
	gfxScale  int
	gfxTitle  string

	OnStart = func() { log.Println("GFX Event: Start") }
	OnEnd   = func() { log.Println("GFX Event: End") }
)

func Init(width, height int, title string, scale int) error {
	gfxWidth = width
	gfxHeight = height
	gfxTitle = title
	gfxScale = scale

	return initialize()
}

func Close(err error) {
	closeChannel <- err
}

func RunSafeReader() error {
	lockChannel = make(chan func())
	closeChannel = make(chan error, 1)

	for {
		select {
		case fn := <-lockChannel:
			fn()
		case err := <-closeChannel:
			OnEnd()
			return err
		}
	}

	return nil
}

func Render(f func() error) { //todo pass scene graph
	err := RunSafeFn(func() error {
		if err := f(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	_ = RunSafeFn(func() error {
		postRender()
		return nil
	})
}

func RunSafeFn(f func() error) error {
	if lockChannel == nil {
		return errors.New("Invalid Safe function. (Channel closed)")
	}

	//reminder: sync.Wait has some problem with runtime.LockOSThread
	w := make(chan error, 1) //todo sync.Pool this

	lockChannel <- func() {
		w <- f()
	}

	return <-w
}
