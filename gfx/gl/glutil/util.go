// Created by nazarigonzalez on 1/1/17.

package glutil

import (
	"errors"

	"github.com/nazariglez/prime/gfx/gl"
)

func CreateProgram(ctx *gl.Context, vertex, fragment string) (*gl.Program, error) {
	program := ctx.CreateProgram()

	vertexShader, err := loadShader(ctx, ctx.VERTEX_SHADER, vertex)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := loadShader(ctx, ctx.FRAGMENT_SHADER, fragment)
	if err != nil {
		return nil, err
	}

	ctx.AttachShader(program, vertexShader)
	ctx.AttachShader(program, fragmentShader)
	ctx.LinkProgram(program)

	ctx.DeleteShader(vertexShader)
	ctx.DeleteShader(fragmentShader)

	if !ctx.GetProgramParameterb(program, ctx.LINK_STATUS) {
		defer ctx.DeleteProgram(program)
		return nil, errors.New("GL Program: " + ctx.GetProgramInfoLog(program))
	}

	return program, nil
}

func loadShader(ctx *gl.Context, typ int, source string) (*gl.Shader, error) {
	s := ctx.CreateShader(typ)
	ctx.ShaderSource(s, source)
	ctx.CompileShader(s)

	if !ctx.GetShaderParameterb(s, ctx.COMPILE_STATUS) {
		defer ctx.DeleteShader(s)
		return nil, errors.New("Shader compile:" + ctx.GetShaderInfoLog(s))
	}

	return s, nil
}
