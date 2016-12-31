// +build js

/**
 * Created by nazarigonzalez on 29/12/16.
 */

package gfx

import (
  "log"

  "github.com/gopherjs/gopherjs/js"
  "honnef.co/go/js/dom"

  "prime/gfx/gl"
  "time"
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

var triangleData = []float32{
  -1, -1, 0,
  1, -1, 0,
  0, 1, 0,
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

  js.Global.Set("ctx", ctx) //todo remove this
  ctx.Viewport(0, 0, gfxWidth, gfxHeight)

  program, err := CreateProgram(ctx, vertexShader, fragmentShader)
  if err != nil {
    log.Println(err)
    return nil
  }

  buff := ctx.CreateBuffer()
  ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
  ctx.BufferData(ctx.ARRAY_BUFFER, triangleData, ctx.STATIC_DRAW)

  ctx.ClearColor(gfxBg[0], gfxBg[1], gfxBg[2], gfxBg[3])

  go func(){
    for {
      //ctx.ClearColor(rand.Float32(),rand.Float32(),rand.Float32(),0)

      ctx.Enable(ctx.DEPTH_TEST)
      ctx.Clear(ctx.COLOR_BUFFER_BIT | ctx.DEPTH_BUFFER_BIT)
      ctx.UseProgram(program)

      ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
      ctx.EnableVertexAttribArray(0)
      ctx.VertexAttribPointer(0, 3, ctx.FLOAT, false, 0, 0)
      ctx.DrawArrays(ctx.TRIANGLES, 0, 3)
      ctx.DisableVertexAttribArray(0)

      //println("lel")
      time.Sleep(33*time.Millisecond)
    }
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

var fragmentShader = `
//#version 120 // OpenGL 2.1
//#version 100 // WebGL 1.0

void main(){
  gl_FragColor = vec4(1.0, 1.0, 0.0, 0.0);
}
`

var vertexShader = `
//#version 120 // OpenGL 2.1
//#version 100 // WebGL 1.0

attribute vec3 vertexPosition_modelspace;

void main(){
  gl_Position.xyz = vertexPosition_modelspace;
  gl_Position.w = 1.0;
}
`