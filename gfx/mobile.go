//Created by nazarigonzalez on 29/12/16.

// +build android

package gfx

import (
  "log"
  "runtime"

  "golang.org/x/mobile/app"
  "golang.org/x/mobile/event/lifecycle"
  "golang.org/x/mobile/event/paint"
  "golang.org/x/mobile/event/size"
  "golang.org/x/mobile/event/touch"

  "prime/gfx/gl"
)

func initialize() error {
  log.Println("Mobile initialized")

  run()
  return nil
}

func run() {
  app.Main(func(a app.App){
    runtime.LockOSThread()

    for e := range a.Events() {

      switch e := a.Filter(e).(type) {

      case lifecycle.Event:

        switch e.Crosses(lifecycle.StageVisible) {

        case lifecycle.CrossOn:
          c, err := gl.NewContext(e.DrawContext)
          if err != nil {
            log.Fatal(err)
            break
          }

          gfxContext = c
          OnStart()

          a.Send(paint.Event{})

        case lifecycle.CrossOff:
          OnEnd()

        }

      case size.Event:

      case paint.Event:
        if gfxContext == nil || e.External {
          continue
        }

        OnDraw()
        a.Publish()
        a.Send(paint.Event{})
      case touch.Event:

      }

    }


  })
}