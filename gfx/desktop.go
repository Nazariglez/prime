// +build !js
// +build !android

/**
 * Created by nazarigonzalez on 29/12/16.
 */

package gfx

import (
  "github.com/go-gl/glfw/v3.2/glfw"
  "log"

  "github.com/pkg/errors"

  "prime/gfx/gl"
)

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

var triangleData = []float32{
  -1, -1, 0,
  1, -1, 0,
  0, 1, 0,
}

func initialize() error {
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

  //ctx.Viewport(0, 0, gfxWidth, gfxHeight) //todo retina issues
  program, err := CreateProgram(ctx, vertexShader, fragmentShader)
  if err != nil {
    return err
  }

  buff := ctx.CreateBuffer()
  ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
  ctx.BufferData(ctx.ARRAY_BUFFER, triangleData, ctx.STATIC_DRAW)

  ctx.ClearColor(gfxBg[0], gfxBg[1], gfxBg[2], gfxBg[3])


  //ctx.BindAttribLocation(program, 0, ctx.Str("vertexPosition_modelspace\x00")) //todo

  for !window.ShouldClose() {
    //draw here

    ctx.Clear(ctx.COLOR_BUFFER_BIT | ctx.DEPTH_BUFFER_BIT)
    ctx.UseProgram(program)

    ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
    ctx.EnableVertexAttribArray(0)
    ctx.VertexAttribPointer(0, 3, ctx.FLOAT, false, 0, 0)
    ctx.DrawArrays(ctx.TRIANGLES, 0, 3)
    ctx.DisableVertexAttribArray(0)

    window.SwapBuffers()
    glfw.PollEvents()
  }

  return nil
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