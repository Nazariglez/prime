// Created by nazarigonzalez on 1/1/17.

package prime

import (
  "log"

  "prime/gfx"
  "prime/gfx/gl"
  "prime/gfx/gl/glutil"
)

type Prime struct {
  ctx *gl.Context
  options *PrimeOptions
}

func runPrime(opts *PrimeOptions) (*Prime, error) {
  engine := &Prime{}
  gfx.OnStart = engine.onGfxStart
  gfx.OnDraw = engine.onGfxDraw
  gfx.OnEnd = engine.onGfxEnd

  engine.options = opts

  if err := gfx.Init(opts.Width, opts.Height, opts.Title, opts.Background[:]); err != nil {
    return nil, err
  }

  return engine, nil
}



var program *gl.Program
var buff *gl.Buffer

func (p *Prime) onGfxStart() {
  ctx, err := gfx.GetContext()
  if err != nil {
    log.Fatal(err)
    return
  }

  p.ctx = ctx
  log.Println("GFX Event: Start")


  //todo remove
  program, err = glutil.CreateProgram(ctx, vertexShader, fragmentShader)
  if err != nil {
    log.Fatal(err)
  }

  buff = ctx.CreateBuffer()
  ctx.BindBuffer(ctx.ARRAY_BUFFER, buff)
  ctx.BufferData(ctx.ARRAY_BUFFER, triangleData, ctx.STATIC_DRAW)

  ctx.ClearColor(
    p.options.Background[0],
    p.options.Background[1],
    p.options.Background[2],
    p.options.Background[3],
  )
}

func (p *Prime) onGfxDraw() {
  log.Println("GFX Event: Draw")


  //todo remove
  p.ctx.Clear(p.ctx.COLOR_BUFFER_BIT | p.ctx.DEPTH_BUFFER_BIT)
  p.ctx.UseProgram(program)

  p.ctx.BindBuffer(p.ctx.ARRAY_BUFFER, buff)
  p.ctx.EnableVertexAttribArray(0)
  p.ctx.VertexAttribPointer(0, 3, p.ctx.FLOAT, false, 0, 0)
  p.ctx.DrawArrays(p.ctx.TRIANGLES, 0, 3)
  p.ctx.DisableVertexAttribArray(0)
}

func (p *Prime) onGfxEnd() {
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