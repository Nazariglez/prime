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
		gfx.GLContext.Clear(gfx.GLContext.COLOR_BUFFER_BIT | gfx.GLContext.DEPTH_BUFFER_BIT)
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
	/*gfx.GLContext.ClearColor(
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	)*/

	gfx.GLContext.UseProgram(triangleProgram)

	gfx.GLContext.BindBuffer(gfx.GLContext.ARRAY_BUFFER, triangleBuffer)
	gfx.GLContext.EnableVertexAttribArray(0)
	gfx.GLContext.VertexAttribPointer(0, 3, gfx.GLContext.FLOAT, false, 0, 0)
	gfx.GLContext.DrawArrays(gfx.GLContext.TRIANGLES, 0, len(triangleData)/3)
	gfx.GLContext.DisableVertexAttribArray(0)

	return nil
}

func GenerateTexture(file string) (*gfx.Texture, error) {
	t := gfx.GLContext.CreateTexture()
	gfx.GLContext.BindTexture(gfx.GLContext.TEXTURE_2D, t)
	/*gfx.GLContext.TexImage2D(
		gfx.GLContext.TEXTURE_2D,
		0,
		gfx.GLContext.RGBA,
		1, 1, 0,
		gfx.GLContext.RGBA,
		gfx.GLContext.UNSIGNED_BYTE,

	)*/

	gfx.GLContext.TexParameteri(gfx.GLContext.TEXTURE_2D, gfx.GLContext.TEXTURE_WRAP_S, gfx.GLContext.CLAMP_TO_EDGE)
	gfx.GLContext.TexParameteri(gfx.GLContext.TEXTURE_2D, gfx.GLContext.TEXTURE_WRAP_T, gfx.GLContext.CLAMP_TO_EDGE)
	gfx.GLContext.TexParameteri(gfx.GLContext.TEXTURE_2D, gfx.GLContext.TEXTURE_MIN_FILTER, gfx.GLContext.LINEAR)
	gfx.GLContext.TexParameteri(gfx.GLContext.TEXTURE_2D, gfx.GLContext.TEXTURE_MAG_FILTER, gfx.GLContext.LINEAR)

	img, err := loadImage(file)
	if err != nil {
		return nil, err
	}

	rect := img.Bounds()
	tex := &gfx.Texture{rect.Dx(), rect.Dy(), t}
	gfx.GLContext.BindTexture(gfx.GLContext.TEXTURE_2D, tex.Tex)
	gfx.GLContext.TexImage2D(gfx.GLContext.TEXTURE_2D, 0, gfx.GLContext.RGBA, gfx.GLContext.RGBA, gfx.GLContext.UNSIGNED_BYTE, img)

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
	texProg, err = glutil.CreateProgram(gfx.GLContext, vst, fst)
	if err != nil {
		return err
	}

	positionLocation = gfx.GLContext.GetAttribLocation(texProg, "a_position")
	texcoordLocation = gfx.GLContext.GetAttribLocation(texProg, "a_texcoord")
	textureLocation = gfx.GLContext.GetUniformLocation(texProg, "u_texture")

	positionBuffer = gfx.GLContext.CreateBuffer()
	gfx.GLContext.BindBuffer(gfx.GLContext.ARRAY_BUFFER, positionBuffer)
	gfx.GLContext.BufferData(gfx.GLContext.ARRAY_BUFFER, positions, gfx.GLContext.STATIC_DRAW)

	texcoordBuffer = gfx.GLContext.CreateBuffer()
	gfx.GLContext.BindBuffer(gfx.GLContext.ARRAY_BUFFER, texcoordBuffer)
	gfx.GLContext.BufferData(gfx.GLContext.ARRAY_BUFFER, texcoords, gfx.GLContext.STATIC_DRAW)


	return nil
}

func DrawTex(tex *gfx.Texture) {
	gfx.GLContext.BindTexture(gfx.GLContext.TEXTURE_2D, tex.Tex)
	gfx.GLContext.UseProgram(texProg)

	gfx.GLContext.BindBuffer(gfx.GLContext.ARRAY_BUFFER, positionBuffer)
	gfx.GLContext.EnableVertexAttribArray(positionLocation)
	gfx.GLContext.VertexAttribPointer(positionLocation, 2, gfx.GLContext.FLOAT, false, 0, 0)
	gfx.GLContext.BindBuffer(gfx.GLContext.ARRAY_BUFFER, colorBuffer)
	gfx.GLContext.EnableVertexAttribArray(texcoordLocation)
	gfx.GLContext.VertexAttribPointer(texcoordLocation, 2, gfx.GLContext.FLOAT, false, 0, 0)

	//camera?

	gfx.GLContext.Uniform1i(textureLocation, 0)
	gfx.GLContext.DrawArrays(gfx.GLContext.TRIANGLES, 0, 6)
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




