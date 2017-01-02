// Created by nazarigonzalez on 1/1/17.

package prime

import (
	"log"

	"prime/gfx"
	"prime/gfx/gl"
	"prime/gfx/gl/glutil"

	"errors"
	"sync"
)

var GL *gl.Context
var CurrentOpts *PrimeOptions
var safeFunc chan func()

func runEngine(opts *PrimeOptions) error {
	CurrentOpts = opts

	gfx.OnStart = onGfxStart
	gfx.OnDraw = onGfxDraw
	gfx.OnEnd = onGfxEnd

	safeFunc = make(chan func())

	if err := gfx.Init(opts.Width, opts.Height, opts.Title, opts.BrowserScale); err != nil {
		return err
	}

	return nil
}

func runOnLockThread(f func() error) error {
	//todo read the safeFunc in lockOsTrehad
	if safeFunc == nil {
		return errors.New("Invalid Safe function. (Channel closed)")
	}

	var wg sync.WaitGroup
	var err error
	wg.Add(1)

	safeFunc <- func() {
		err = f()
		wg.Done()
	}

	wg.Wait()
	return err
}

var program *gl.Program
var buff *gl.Buffer

func onGfxStart() {
	go func() {
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
			CurrentOpts.Background[0],
			CurrentOpts.Background[1],
			CurrentOpts.Background[2],
			CurrentOpts.Background[3],
		)
	}()
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
