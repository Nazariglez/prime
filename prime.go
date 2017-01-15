// Created by nazarigonzalez on 1/1/17.

package prime

import (
	"log"

	"prime/gfx"
	"prime/gfx/gl"
	"prime/gfx/gl/glutil"

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

var tex *gfx.Texture

func onGfxStart() {
	log.Println("GFX Event: Start")

	err := gfx.RunSafeFn(func() error {
		InitTex()
		t, err := GenerateTexture("./texture.png")
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
//precision mediump float;

varying vec2 v_texcoord;

uniform sampler2D u_texture;

void main() {
	gl_FragColor = texture2D(u_texture, v_texcoord);
}

`

var vst = `
attribute vec4 a_position;
attribute vec2 a_texcoord;

uniform mat4 u_matrix;

varying vec2 v_texcoord;

void main() {
   gl_Position = u_matrix * a_position;
   v_texcoord = a_texcoord;
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
	/*gfx.GLContext.ClearColor(
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	)*/

	gfx.GL.UseProgram(triangleProgram)

	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, triangleBuffer)
	gfx.GL.EnableVertexAttribArray(0)
	gfx.GL.VertexAttribPointer(0, 3, gfx.GL.FLOAT, false, 0, 0)
	gfx.GL.DrawArrays(gfx.GL.TRIANGLES, 0, len(triangleData)/3)
	gfx.GL.DisableVertexAttribArray(0)

	return nil
}

func GenerateTexture(file string) (*gfx.Texture, error) {
	t := gfx.GL.CreateTexture()
	gfx.GL.BindTexture(gfx.GL.TEXTURE_2D, t)
	/*gfx.GLContext.TexImage2D(
		gfx.GLContext.TEXTURE_2D,
		0,
		gfx.GLContext.RGBA,
		1, 1, 0,
		gfx.GLContext.RGBA,
		gfx.GLContext.UNSIGNED_BYTE,

	)*/

	gfx.GL.TexParameteri(gfx.GL.TEXTURE_2D, gfx.GL.TEXTURE_WRAP_S, gfx.GL.CLAMP_TO_EDGE)
	gfx.GL.TexParameteri(gfx.GL.TEXTURE_2D, gfx.GL.TEXTURE_WRAP_T, gfx.GL.CLAMP_TO_EDGE)
	gfx.GL.TexParameteri(gfx.GL.TEXTURE_2D, gfx.GL.TEXTURE_MIN_FILTER, gfx.GL.LINEAR)
	gfx.GL.TexParameteri(gfx.GL.TEXTURE_2D, gfx.GL.TEXTURE_MAG_FILTER, gfx.GL.LINEAR)

	img, err := loadImage(file)
	if err != nil {
		return nil, err
	}

	rect := img.Bounds()
	tex := &gfx.Texture{rect.Dx(), rect.Dy(), t}
	gfx.GL.BindTexture(gfx.GL.TEXTURE_2D, tex.Tex)
	gfx.GL.TexImage2D(gfx.GL.TEXTURE_2D, 0, gfx.GL.RGBA, gfx.GL.RGBA, gfx.GL.UNSIGNED_BYTE, img)

	return tex, nil
}

var texProg *gl.Program
var positionBuffer *gl.Buffer
var texcoordBuffer *gl.Buffer
var positionLocation int
var colorBuffer *gl.Buffer
var texcoordLocation int
var textureLocation *gl.UniformLocation

var positions = []float32{
	0, 0, 0,
	1, 1, 0,
	1, 0, 0,
	1, 1, 1,
}

var texcoords = []float32{
	0, 0, 0,
	1, 1, 0,
	1, 0, 0,
	1, 1, 1,
}

func InitTex() error {
	var err error
	texProg, err = glutil.CreateProgram(gfx.GL, vst, fst)
	if err != nil {
		return err
	}

	positionLocation = gfx.GL.GetAttribLocation(texProg, "a_position")
	texcoordLocation = gfx.GL.GetAttribLocation(texProg, "a_texcoord")
	textureLocation = gfx.GL.GetUniformLocation(texProg, "u_texture")

	positionBuffer = gfx.GL.CreateBuffer()
	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, positionBuffer)
	gfx.GL.BufferData(gfx.GL.ARRAY_BUFFER, positions, gfx.GL.STATIC_DRAW)

	texcoordBuffer = gfx.GL.CreateBuffer()
	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, texcoordBuffer)
	gfx.GL.BufferData(gfx.GL.ARRAY_BUFFER, texcoords, gfx.GL.STATIC_DRAW)


	return nil
}

func DrawTex(tex *gfx.Texture) {
	gfx.GL.BindTexture(gfx.GL.TEXTURE_2D, tex.Tex)
	gfx.GL.UseProgram(texProg)

	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, positionBuffer)
	gfx.GL.EnableVertexAttribArray(positionLocation)
	gfx.GL.VertexAttribPointer(positionLocation, 2, gfx.GL.FLOAT, false, 0, 0)
	gfx.GL.BindBuffer(gfx.GL.ARRAY_BUFFER, colorBuffer)
	gfx.GL.EnableVertexAttribArray(texcoordLocation)
	gfx.GL.VertexAttribPointer(texcoordLocation, 2, gfx.GL.FLOAT, false, 0, 0)

	//camera?

	gfx.GL.Uniform1i(textureLocation, 0)
	gfx.GL.DrawArrays(gfx.GL.TRIANGLES, 0, 4)

	//todo http://webglfundamentals.org/webgl/lessons/webgl-2d-drawimage.html
	//todo 2D image view-source:http://webglfundamentals.org/webgl/webgl-2d-image.html
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




