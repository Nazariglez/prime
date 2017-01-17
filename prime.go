// Created by nazarigonzalez on 1/1/17.

package prime

import (
	"log"

	"prime/gfx"
	"prime/gfx/gl"
	"prime/gfx/gl/glutil"

	"math/rand"
	"prime/assets"
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

var tex *gfx.Texture

func onGfxStart() {
	log.Println("GFX Event: Start")

	err := gfx.RunSafeFn(func() error {
		if err := InitTex(); err != nil {
			return err
		}

		t, err := GenerateTexture("texture.png")
		if err != nil {
			return err
		}

		tex = t
		return nil
	})
	//err := gfx.RunSafeFn(drawTriangleInit)
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
		gfx.GL.Clear(gfx.GL.COLOR_BUFFER_BIT | gfx.GL.DEPTH_BUFFER_BIT)
		DrawTex(tex)
		return nil
	})
	//gfx.Render(drawTriangleRender)
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

/*var triangleData = []float32{
	-1, -1, 0,
	1, -1, 0,
	0, 1, 0,

	1, 1, 0,
	-1, 1, 0,
	0, -1, 0,
}*/

var fst = `
precision mediump float;

// our texture
uniform sampler2D u_image;

// the texCoords passed in from the vertex shader.
varying vec2 v_texCoord;

void main() {
   gl_FragColor = texture2D(u_image, v_texCoord);
}

`

var vst = `
attribute vec2 a_position;
attribute vec2 a_texCoord;

uniform vec2 u_resolution;

varying vec2 v_texCoord;

void main() {
   // convert the rectangle from pixels to 0.0 to 1.0
   vec2 zeroToOne = a_position / u_resolution;

   // convert from 0->1 to 0->2
   vec2 zeroToTwo = zeroToOne * 2.0;

   // convert from 0->2 to -1->+1 (clipspace)
   vec2 clipSpace = zeroToTwo - 1.0;

   gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);

   // pass the texCoord to the fragment shader
   // The GPU will interpolate this value between points.
   v_texCoord = a_texCoord;
}
`

var triangleData = []float32{
	-1, -1, 0,
	1, -1, 0,
	-1, 1, 0,

	-1, 1, 0,
	1, -1, 0,
	1, 1, 0,
}

func drawTriangleInit() error {
	p, err := glutil.CreateProgram(gfx.GL, triangleVertexShader, triangleFragmentShader)
	if err != nil {
		return err
	}

	triangleProgram = p

	triangleBuffer = gfx.GL.CreateBuffer()
	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, triangleBuffer)
	gfx.GL.BufferData(gfx.GL.ARRAY_BUFFER, triangleData, gfx.GL.STATIC_DRAW)

	gfx.GL.ClearColor(
		CurrentOpts.Background[0],
		CurrentOpts.Background[1],
		CurrentOpts.Background[2],
		CurrentOpts.Background[3],
	)
	return nil
}

func drawTriangleRender() error {
	gfx.GL.Clear(gfx.GL.COLOR_BUFFER_BIT | gfx.GL.DEPTH_BUFFER_BIT)
	gfx.GL.ClearColor(
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	)

	gfx.GL.UseProgram(triangleProgram)

	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, triangleBuffer)
	gfx.GL.EnableVertexAttribArray(0)
	gfx.GL.VertexAttribPointer(0, 3, gfx.GL.FLOAT, false, 0, 0)
	gfx.GL.DrawArrays(gfx.GL.TRIANGLES, 0, len(triangleData)/3)
	gfx.GL.DisableVertexAttribArray(0)

	return nil
}

func GenerateTexture(file string) (*gfx.Texture, error) {
	img, err := assets.LoadImage(file)
	if err != nil {
		return nil, err
	}

	return gfx.NewTexture(img), nil
}

var texProg *gl.Program
var positionBuffer *gl.Buffer
var texcoordBuffer *gl.Buffer
var positionLocation int
var texcoordLocation int
var textureLocation *gl.UniformLocation
var resolutionLocation *gl.UniformLocation

var positions = []float32{
	0, 0, 0,
	1, 1, 0,
	1, 0, 0,
	1, 1, 1,
}

var texcoords = []float32{
	0, 0,
	1, 0,
	0, 1,
	0, 1,
	1, 0,
	1, 1,
}

func InitTex() error {
	var err error
	texProg, err = glutil.CreateProgram(gfx.GL, vst, fst)
	if err != nil {
		return err
	}

	positionLocation = gfx.GL.GetAttribLocation(texProg, "a_position")
	texcoordLocation = gfx.GL.GetAttribLocation(texProg, "a_texCoord")

	positionBuffer = gfx.GL.CreateBuffer()
	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, positionBuffer)
	setRectangle(0, 0, 256, 256)

	//gfx.GL.BufferData(gfx.GL.ARRAY_BUFFER, positions, gfx.GL.STATIC_DRAW)

	texcoordBuffer = gfx.GL.CreateBuffer()
	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, texcoordBuffer)
	gfx.GL.BufferData(gfx.GL.ARRAY_BUFFER, texcoords, gfx.GL.STATIC_DRAW)

	resolutionLocation = gfx.GL.GetUniformLocation(texProg, "u_resolution")

	//todo viewport
	return nil
}

func setRectangle(x, y, width, height int) {
	x1 := float32(x)
	x2 := float32(x + width)
	y1 := float32(y)
	y2 := float32(y + height)

	gfx.GL.BufferData(gfx.GL.ARRAY_BUFFER, []float32{
		x1, y1,
		x2, y1,
		x1, y2,
		x1, y2,
		x2, y1,
		x2, y2,
	}, gfx.GL.STATIC_DRAW)
}

var s float32 = 255
var pingPong = true

func DrawTex(tex *gfx.Texture) {
	gfx.GL.Clear(gfx.GL.COLOR_BUFFER_BIT | gfx.GL.DEPTH_BUFFER_BIT)
	gfx.GL.ClearColor(
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	)

	//gfx.GL.BindTexture(gfx.GL.TEXTURE_2D, tex.Tex)
	gfx.GL.UseProgram(texProg)

	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, positionBuffer)
	gfx.GL.EnableVertexAttribArray(positionLocation)
	gfx.GL.VertexAttribPointer(positionLocation, 2, gfx.GL.FLOAT, false, 0, 0)
	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, texcoordBuffer)
	gfx.GL.EnableVertexAttribArray(texcoordLocation)
	gfx.GL.VertexAttribPointer(texcoordLocation, 2, gfx.GL.FLOAT, false, 0, 0)

	//camera?

	if pingPong {
		if s > 800 {
			pingPong = false
		}
		s++
	} else {
		if s < 256 {
			pingPong = true
		}
		s--
	}

	gfx.GL.Uniform2f(resolutionLocation, s, s)
	gfx.GL.DrawArrays(gfx.GL.TRIANGLES, 0, 6)
	//gfx.GL.DrawElements(gfx.GL.TRIANGLES, 4, gfx.GL.UNSIGNED_SHORT, 0)
	//gfx.GL.DisableVertexAttribArray(positionLocation)
	//gfx.GL.DisableVertexAttribArray(texcoordLocation)

	//todo http://webglfundamentals.org/webgl/lessons/webgl-2d-drawimage.html
	//todo 2D image view-source:http://webglfundamentals.org/webgl/webgl-2d-image.html
}
