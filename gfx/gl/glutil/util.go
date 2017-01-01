// Created by nazarigonzalez on 1/1/17.

package glutil

import (
  "errors"

  "prime/gfx/gl"
)

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
