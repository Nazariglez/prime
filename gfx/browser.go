// Created by nazarigonzalez on 29/12/16.

// +build js

package gfx

import (
  "log"
  "time"
  "runtime"

  "github.com/gopherjs/gopherjs/js"
  "honnef.co/go/js/dom"

  "prime/gfx/gl"
)


type Texture struct{ *js.Object }
type Buffer struct{ *js.Object }
type FrameBuffer struct{ *js.Object }
type RenderBuffer struct{ *js.Object }
type Program struct{ *js.Object }
type UniformLocation struct{ *js.Object }
type Shader struct{ *js.Object }

var htmlContentLoaded bool


func initialize() error {
  runtime.LockOSThread()
  log.Println("Browser initialized")

  doc := dom.GetWindow().Document().(dom.HTMLDocument)
  doc.SetTitle(gfxTitle)

  onReady(func(){
    log.Println("Document Loaded!!")

    var canvas *dom.HTMLCanvasElement
    if doc.GetElementByID("prime-view") == nil {
      canvas = doc.CreateElement("canvas").(*dom.HTMLCanvasElement)
      canvas.Set("id", "prime-view")
      doc.Body().AppendChild(canvas)
    } else {
      canvas = doc.GetElementByID("prime-view").(*dom.HTMLCanvasElement)
    }

    canvas.Width = gfxWidth
    canvas.Height = gfxHeight

    if err := run(canvas); err != nil {
      log.Println(err)
    }
  })

  return nil
}

func run(canvas *dom.HTMLCanvasElement) error {
  attrs := gl.DefaultAttributes()
  attrs.Alpha = false
  attrs.Depth = false
  attrs.PremultipliedAlpha = false
  attrs.PreserveDrawingBuffer = false
  attrs.Antialias = false

  ctx, err := gl.NewContext(canvas.Object, attrs)
  if err != nil {
    return err
  }
  gfxContext = ctx

  OnStart()

  js.Global.Set("ctx", ctx) //todo remove this
  ctx.Viewport(0, 0, gfxWidth, gfxHeight) //todo fix this, retina issues (removed from here)

  go func(){
    for {
      OnDraw()
      time.Sleep(33*time.Millisecond)
    }

    OnEnd()
  }()

  return nil
}

func onReady(cb func()) {
  d := js.Global.Get("document")

  if isReadyStateComplete() {
    htmlContentLoaded = true
    cb()
    return
  }

  if d.Get("addEventListener") != nil {
    d.Call("addEventListener", "DOMContentLoaded", onLoad(cb), false)
    js.Global.Call("addEventListener", "load", onLoad(cb), false)
  } else {
    d.Call("attachEvent", "onreadystatechange", func(){
      if isReadyStateComplete() {
        htmlContentLoaded = true
        cb()
      }
    })
    js.Global.Call("attachEvent", "onload", onLoad(cb))
  }

}

func onLoad(cb func()) func() {
  return func(){
    if htmlContentLoaded {
      return
    }

    htmlContentLoaded = true
    cb()
  }
}

func isReadyStateComplete() bool {
  return js.Global.Get("document").Get("readyState").String() == "complete"
}