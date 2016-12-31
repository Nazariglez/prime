// +build android

/**
 * Created by nazarigonzalez on 29/12/16.
 */

package gfx

import (
  "log"
  "errors"

  "golang.org/x/mobile/app"
  "golang.org/x/mobile/event/lifecycle"
  "golang.org/x/mobile/event/paint"
  "golang.org/x/mobile/event/size"
  "golang.org/x/mobile/event/touch"

  "prime/gfx/gl"
)

var (
  program  *gl.Program
  buff *gl.Buffer
)

func initialize() error {
  log.Println("Mobile initialized")

  run()
  return nil
}

func run() {
  app.Main(func(a app.App){
    var ctx *gl.Context
    var sz size.Event

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

          ctx = c
          onStart(ctx)
          a.Send(paint.Event{})

        case lifecycle.CrossOff:
          onStop(ctx)
          ctx = nil

        }

      case size.Event:
        sz = e
      case paint.Event:
        if ctx == nil || e.External {
          continue
        }

        onPaint(ctx, sz)
        a.Publish()
        a.Send(paint.Event{})
      case touch.Event:

      }

    }


  })
}

func onStart(ctx *gl.Context) {
  var err error
  program, err = CreateProgram(ctx, vertexShader, fragmentShader)
  if err != nil {
    log.Printf("error creating GL program: %v", err)
    return
  }

  t := []float32{
    -1, -1, 0,
    1, -1, 0,
    0, 1, 0,
  }

  buff = ctx.CreateBuffer()
  ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
  ctx.BufferData(ctx.ARRAY_BUFFER, t, ctx.STATIC_DRAW)

  ctx.ClearColor(gfxBg[0], gfxBg[1], gfxBg[2], gfxBg[3])
}

func onStop(ctx *gl.Context) {
  ctx.DeleteProgram(program)
  ctx.DeleteBuffer(buff)
}

func onPaint(ctx *gl.Context, sz size.Event) {
  ctx.Clear(ctx.COLOR_BUFFER_BIT | ctx.DEPTH_BUFFER_BIT)
  ctx.UseProgram(program)

  ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
  ctx.EnableVertexAttribArray(0)
  ctx.VertexAttribPointer(0, 3, ctx.FLOAT, false, 0, 0)
  ctx.DrawArrays(ctx.TRIANGLES, 0, 3)
  ctx.DisableVertexAttribArray(0)

  log.Println("PAINT PAINT PAINT")
}

func CreateProgram(ctx *gl.Context, v, f string) (*gl.Program, error) {
  program := ctx.CreateProgram()

  vertexShader := ctx.CreateShader(ctx.VERTEX_SHADER)
  ctx.ShaderSource(vertexShader, v)
  ctx.CompileShader(vertexShader)

  if !ctx.GetShaderParameterb(vertexShader, ctx.COMPILE_STATUS) {
    defer ctx.DeleteShader(vertexShader)
    return &gl.Program{}, errors.New("Shader compile: " + ctx.GetShaderInfoLog(vertexShader))
  }

  fragmentShader := ctx.CreateShader(ctx.FRAGMENT_SHADER)
  ctx.ShaderSource(fragmentShader, f)
  ctx.CompileShader(fragmentShader)

  if !ctx.GetShaderParameterb(fragmentShader, ctx.COMPILE_STATUS) {
    defer ctx.DeleteShader(fragmentShader)
    return &gl.Program{}, errors.New("Shader compile: " + ctx.GetShaderInfoLog(fragmentShader))
  }

  ctx.AttachShader(program, vertexShader)
  ctx.AttachShader(program, fragmentShader)
  ctx.LinkProgram(program)

  ctx.DeleteShader(vertexShader)
  ctx.DeleteShader(fragmentShader)

  if !ctx.GetProgramParameterb(program, ctx.LINK_STATUS) {
    defer ctx.DeleteProgram(program)
    return &gl.Program{}, errors.New("GL Program: " + ctx.GetProgramInfoLog(program))
  }

  return program, nil
}

var fragmentShader = `
//#version 120 // OpenGL 2.1

void main(){
  gl_FragColor = vec4(1.0, 1.0, 0.0, 0.0);
}
`

var vertexShader = `
//#version 120 // OpenGL 2.1

attribute vec3 vertexPosition_modelspace;

void main(){
  gl_Position.xyz = vertexPosition_modelspace;
  gl_Position.w = 1.0;
}
`