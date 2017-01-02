// Created by nazarigonzalez on 1/1/17.

package prime

import (
	"log"

	"prime/gfx"
	"prime/gfx/gl"
	"prime/gfx/gl/glutil"
)

var GL *gl.Context
var currentOptions *PrimeOptions

func runEngine(opts *PrimeOptions) error {
	currentOptions = opts

	gfx.OnStart = onGfxStart
	gfx.OnDraw = onGfxDraw
	gfx.OnEnd = onGfxEnd

	if err := gfx.Init(opts.Width, opts.Height, opts.Title); err != nil {
		return err
	}

	return nil
}

var program *gl.Program
var buff *gl.Buffer

func onGfxStart() {
	ctx, err := gfx.GetContext()
	if err != nil {
		log.Fatal(err)
		return
	}

	GL = ctx
	log.Println("GFX Event: Start")

	//todo remove
	program, err = glutil.CreateProgram(GL, vertexShader, fragmentShader)
	if err != nil {
		log.Fatal(err)
	}

	buff = GL.CreateBuffer()
	GL.BindBuffer(GL.ARRAY_BUFFER, buff)
	GL.BufferData(GL.ARRAY_BUFFER, triangleData, GL.STATIC_DRAW)

	GL.ClearColor(
		currentOptions.Background[0],
		currentOptions.Background[1],
		currentOptions.Background[2],
		currentOptions.Background[3],
	)
}

func onGfxDraw() {
	log.Println("GFX Event: Draw")

	//todo remove
	GL.Clear(GL.COLOR_BUFFER_BIT | GL.DEPTH_BUFFER_BIT)
	GL.UseProgram(program)

	GL.BindBuffer(GL.ARRAY_BUFFER, buff)
	GL.EnableVertexAttribArray(0)
	GL.VertexAttribPointer(0, 3, GL.FLOAT, false, 0, 0)
	GL.DrawArrays(GL.TRIANGLES, 0, 3)
	GL.DisableVertexAttribArray(0)
}

func onGfxEnd() {
	log.Println("GFX Event: End")

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

var triangleData = []float32{
	-1, -1, 0,
	1, -1, 0,
	0, 1, 0,
}
