// Created by nazarigonzalez on 1/1/17.

package prime

import (
	"log"

	"prime/gfx"
	"prime/gfx/gl"
	"prime/gfx/gl/glutil"

	"time"
	"math/rand"
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

	var err error
	go gfx.RunSafeFn(func() error {
		program, err = glutil.CreateProgram(ctx, vertexShader, fragmentShader)
		if err != nil {
			log.Fatal(err)
		}

		buff = ctx.CreateBuffer()
		ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
		ctx.BufferData(ctx.ARRAY_BUFFER, triangleData, ctx.STATIC_DRAW)

		ctx.ClearColor(
			CurrentOpts.Background[0],
			CurrentOpts.Background[1],
			CurrentOpts.Background[2],
			CurrentOpts.Background[3],
		)
		log.Println("OpenOpen")
		return nil
	})

	go func(){
		for {
			go gfx.Render(func() error {
				ctx.Clear(ctx.COLOR_BUFFER_BIT | ctx.DEPTH_BUFFER_BIT)
				ctx.ClearColor(
					rand.Float32(),
					rand.Float32(),
					rand.Float32(),
					rand.Float32(),
				)

				ctx.UseProgram(program)

				ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
				ctx.EnableVertexAttribArray(0)
				ctx.VertexAttribPointer(0, 3, ctx.FLOAT, false, 0, 0)
				ctx.DrawArrays(ctx.TRIANGLES, 0, 3)
				ctx.DisableVertexAttribArray(0)

				return nil
			})
			time.Sleep(16 * time.Millisecond)
		}
	}()
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
