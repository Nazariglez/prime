// Created by nazarigonzalez on 1/1/17.

package glutil

import (
	"errors"
	"strings"
	"strconv"

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

var fragTestIfs = `
#ifdef GL_ES
precision mediump float;
#endif

void main(){
	float test = 0.1;
	%loop%
	gl_FragColor = vec4(0.0);
}
`

func GetMaxTextures(ctx *gl.Context) int {
	max := ctx.GetParameter(ctx.MAX_TEXTURE_IMAGE_UNITS)
	shader := ctx.CreateShader(ctx.FRAGMENT_SHADER)
	defer ctx.DeleteShader(shader)

	for {
		f := strings.Replace(fragTestIfs, "%loop%", generateIfsTest(max), -1)
		ctx.ShaderSource(shader, f)
		ctx.CompileShader(shader)
		if !ctx.GetShaderParameterb(shader, ctx.COMPILE_STATUS) {
			max /= 2
		} else {
			break
		}
	}

	return max
}

func generateIfsTest(max int) string {
	str := ""
	for i := 0; i < max; i++ {
		if i > 0 {
			str += "\nelse "
		}

		if i < max - 1 {
			str += "if(test == " + strconv.Itoa(i) + ".0){}"
		}
	}

	return str
}