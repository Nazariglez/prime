//Created by nazarigonzalez on 29/12/16.

// +build android

package gfx

import (
	"log"
	"runtime"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"

	"prime/gfx/gl"
)

var mobileApp app.App
var mobileStop bool

func initialize() error {
	log.Println("Mobile initialized")
	return run()
}

func run() error {
	ch := make(chan error)

	app.Main(func(a app.App) {
		runtime.LockOSThread()
		mobileApp = a

		w := make(chan error)

		go func(){
			for e := range a.Events() {

				switch e := a.Filter(e).(type) {

				case lifecycle.Event:

					switch e.Crosses(lifecycle.StageVisible) {

					case lifecycle.CrossOn:
						c, err := gl.NewContext(e.DrawContext)
						if err != nil {
							w <- err
							break
						}

						GLContext = c
						OnStart()

						w <- nil

					case lifecycle.CrossOff:
						OnEnd()

					}

				case size.Event:

				/*case paint.Event:
					if GLContext == nil || e.External {
						continue
					}

						select {
						case fn := <-lockChannel:
							fn()
						}

					a.Publish()
					a.Send(paint.Event{})*/
				case touch.Event:

				}

			}
		}()

		if err := <- w; err != nil {
			ch <- err
			return
		}
		close(w)

		for !mobileStop {
			select {
			case fn := <-lockChannel:
				fn()
			}
		}

	})

	return <-ch
}

func postRender() {
	mobileApp.Publish()
}