/**
 * Created by nazarigonzalez on 29/12/16.
 */

package gfx

import (
  "runtime"

  "errors"
  "prime/gfx/gl"
)

var (
  gfxWidth int
  gfxHeight int
  gfxTitle string
  gfxBg []float32
)

func Init(width, height int, title string, bg []float32) error {
  runtime.LockOSThread()
  gfxWidth = width
  gfxHeight = height
  gfxTitle = title
  gfxBg = bg

  return initialize()
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

