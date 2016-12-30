// +build android

/**
 * Created by nazarigonzalez on 29/12/16.
 */

package gfx

import (
  "encoding/binary"
  "log"

  "golang.org/x/mobile/app"
  "golang.org/x/mobile/event/lifecycle"
  "golang.org/x/mobile/event/paint"
  "golang.org/x/mobile/event/size"
  "golang.org/x/mobile/event/touch"
  "golang.org/x/mobile/exp/app/debug"
  "golang.org/x/mobile/exp/f32"
  "golang.org/x/mobile/exp/gl/glutil"
  "golang.org/x/mobile/gl"
)

var (
  images   *glutil.Images
  fps      *debug.FPS
  program  gl.Program
  position gl.Attrib
  offset   gl.Uniform
  color    gl.Uniform
  buff gl.Buffer

  green  float32
  touchX float32
  touchY float32
)

func initialize() error {
  log.Println("Mobile initialized")

  run()
  return nil
}

func run() {
  app.Main(func(a app.App) {
    var glctx gl.Context
    var sz size.Event
    for e := range a.Events() {
      switch e := a.Filter(e).(type) {
      case lifecycle.Event:
        switch e.Crosses(lifecycle.StageVisible) {
        case lifecycle.CrossOn:
          glctx, _ = e.DrawContext.(gl.Context)
          onStart(glctx)
          a.Send(paint.Event{})
        case lifecycle.CrossOff:
          onStop(glctx)
          glctx = nil
        }
      case size.Event:
        sz = e
        touchX = float32(sz.WidthPx / 2)
        touchY = float32(sz.HeightPx / 2)
      case paint.Event:
        if glctx == nil || e.External {
          // As we are actively painting as fast as
          // we can (usually 60 FPS), skip any paint
          // events sent by the system.
          continue
        }

        onPaint(glctx, sz)
        a.Publish()
        // Drive the animation by preparing to paint the next frame
        // after this one is shown.
        a.Send(paint.Event{})
      case touch.Event:
        touchX = e.X
        touchY = e.Y
      }
    }
  })
}

func onStart(glctx gl.Context) {
  var err error
  program, err = glutil.CreateProgram(glctx, vertexShader1, fragmentShader1)
  if err != nil {
    log.Printf("error creating GL program: %v", err)
    return
  }

  buff = glctx.CreateBuffer()
  glctx.BindBuffer(gl.ARRAY_BUFFER, buff)
  glctx.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW)

  //position = glctx.GetAttribLocation(program, "position")
  //color = glctx.GetUniformLocation(program, "color")
  //offset = glctx.GetUniformLocation(program, "offset")

  //images = glutil.NewImages(glctx)
  //fps = debug.NewFPS(images)

  glctx.ClearColor(gfxBg[0], gfxBg[1], gfxBg[2], gfxBg[3])
}

func onStop(glctx gl.Context) {
  glctx.DeleteProgram(program)
  glctx.DeleteBuffer(buff)
  //fps.Release()
  //images.Release()
}

func onPaint(glctx gl.Context, sz size.Event) {
  glctx.Clear(gl.COLOR_BUFFER_BIT)

  glctx.UseProgram(program)

  /*green += 0.01
  if green > 1 {
    green = 0
  }
  glctx.Uniform4f(color, 0, green, 0, 1)

  glctx.Uniform2f(offset, touchX/float32(sz.WidthPx), touchY/float32(sz.HeightPx))
*/
  glctx.BindBuffer(gl.ARRAY_BUFFER, buff)
  glctx.EnableVertexAttribArray(position)
  glctx.VertexAttribPointer(position, coordsPerVertex, gl.FLOAT, false, 0, 0)
  glctx.DrawArrays(gl.TRIANGLES, 0, vertexCount)
  glctx.DisableVertexAttribArray(position)

  log.Println("PAINT PAINT PAINT")

  //fps.Draw(sz)
}

var triangleData = f32.Bytes(binary.LittleEndian,
  -1, -1, 0,
  1, -1, 0,
  0, 1, 0,
)

const (
  coordsPerVertex = 3
  vertexCount     = 3
)

const vertexShader = `#version 100
uniform vec2 offset;
attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	gl_Position = position + offset4;
}`

const fragmentShader = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`

var fragmentShader1 = `
//#version 120 // OpenGL 2.1

void main(){
  gl_FragColor = vec4(1.0, 1.0, 0.0, 0.0);
}
`

var vertexShader1 = `
//#version 120 // OpenGL 2.1

attribute vec3 vertexPosition_modelspace;

void main(){
  gl_Position.xyz = vertexPosition_modelspace;
  gl_Position.w = 1.0;
}
`