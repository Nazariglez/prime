// Created by nazarigonzalez on 1/1/17.

package prime

import (
	"log"

	"prime/gfx"
	"prime/gfx/gl"
	"prime/gfx/gl/glutil"

	"math/rand"
	"prime/loop"
)

var CurrentOpts *PrimeOptions

func runEngine(opts *PrimeOptions) error {
	CurrentOpts = opts

	gfx.OnStart = onGfxStart
	gfx.OnEnd = onGfxEnd

	if err := gfx.Init(opts.Width, opts.Height, opts.Title, opts.BrowserScale); err != nil {
		return err
	}

	return nil
}

var program *gl.Program
var buff *gl.Buffer

func onGfxStart() {
	ctx := gfx.GLContext

	log.Println("GFX Event: Start")

	err := gfx.RunSafeFn(func() error {
		p, err := glutil.CreateProgram(ctx, vertexShader, fragmentShader)
		if err != nil {
			return err
		}

		program = p

		buff = ctx.CreateBuffer()
		ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
		ctx.BufferData(ctx.ARRAY_BUFFER, triangleData, ctx.STATIC_DRAW)

		ctx.ClearColor(
			CurrentOpts.Background[0],
			CurrentOpts.Background[1],
			CurrentOpts.Background[2],
			CurrentOpts.Background[3],
		)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	loop.Run(update)

}

func onGfxEnd() {
	log.Println("GFX Event: End")

}

func update(d float64) {
	gfx.Render(func() error {
		gfx.GLContext.Clear(gfx.GLContext.COLOR_BUFFER_BIT | gfx.GLContext.DEPTH_BUFFER_BIT)
		gfx.GLContext.ClearColor(
			rand.Float32(),
			rand.Float32(),
			rand.Float32(),
			rand.Float32(),
		)

		gfx.GLContext.UseProgram(program)

		gfx.GLContext.BindBuffer(gfx.GLContext.ARRAY_BUFFER, buff)
		gfx.GLContext.EnableVertexAttribArray(0)
		gfx.GLContext.VertexAttribPointer(0, 3, gfx.GLContext.FLOAT, false, 0, 0)
		gfx.GLContext.DrawArrays(gfx.GLContext.TRIANGLES, 0, 3)
		gfx.GLContext.DisableVertexAttribArray(0)

		return nil
	})
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
