// Created by nazarigonzalez on 1/1/17.

package prime

import (
	"log"

	"prime/gfx"
	"prime/gfx/gl"
	"prime/gfx/gl/glutil"

	"math/rand"
	"prime/loop"

	"image"
	_"image/png"
	"os"
	"image/draw"
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

func onGfxStart() {
	log.Println("GFX Event: Start")

	err := gfx.RunSafeFn(drawTriangleInit)
	if err != nil {
		log.Fatal(err)
	}

	loop.Run(update)
}

func onGfxEnd() {
	log.Println("GFX Event: End")

}

func update(d float64) {
	gfx.Render(drawTriangleRender)
}


//TRIANGLE
var triangleProgram *gl.Program
var triangleBuffer *gl.Buffer

var triangleFragmentShader = `
//#version 120 // OpenGL 2.1

void main(){
  gl_FragColor = vec4(1.0, 1.0, 0.0, 0.0);
}
`

var triangleVertexShader = `
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

func drawTriangleInit() error {
	p, err := glutil.CreateProgram(gfx.GLContext, triangleVertexShader, triangleFragmentShader)
	if err != nil {
		return err
	}

	triangleProgram = p

	triangleBuffer = gfx.GLContext.CreateBuffer()
	gfx.GLContext.BindBuffer(gfx.GLContext.ARRAY_BUFFER, triangleBuffer)
	gfx.GLContext.BufferData(gfx.GLContext.ARRAY_BUFFER, triangleData, gfx.GLContext.STATIC_DRAW)

	gfx.GLContext.ClearColor(
		CurrentOpts.Background[0],
		CurrentOpts.Background[1],
		CurrentOpts.Background[2],
		CurrentOpts.Background[3],
	)
	return nil
}

func drawTriangleRender() error {
	gfx.GLContext.Clear(gfx.GLContext.COLOR_BUFFER_BIT | gfx.GLContext.DEPTH_BUFFER_BIT)
	gfx.GLContext.ClearColor(
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	)

	gfx.GLContext.UseProgram(triangleProgram)

	gfx.GLContext.BindBuffer(gfx.GLContext.ARRAY_BUFFER, triangleBuffer)
	gfx.GLContext.EnableVertexAttribArray(0)
	gfx.GLContext.VertexAttribPointer(0, 3, gfx.GLContext.FLOAT, false, 0, 0)
	gfx.GLContext.DrawArrays(gfx.GLContext.TRIANGLES, 0, 3)
	gfx.GLContext.DisableVertexAttribArray(0)

	return nil
}

func loadImage(img string) (*image.NRGBA, error) {
	f, err := os.Open(img)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	i, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	rgba := image.NewNRGBA(i.Bounds())
	draw.Draw(rgba, rgba.Bounds(), i, image.Point{0, 0}, draw.Src)

	return rgba, nil
}




