// Created by nazarigonzalez on 25/1/17.

package gfx

import (
  "strings"
  "strconv"
  "log"

  "github.com/nazariglez/prime/gfx/gl"
  "github.com/nazariglez/prime/gfx/gl/glutil"
  //"github.com/gopherjs/gopherjs/js"
)

//https://github.com/pixijs/pixi-gl-core/blob/master/src/GLShader.js
//https://github.com/pixijs/pixi.js/blob/dev/src/core/Shader.js

type Shader struct {
  Program *gl.Program
  UniformData map[string]*ShaderUniform
  Uniforms int
  Attributes map[string]*ShaderAttrib
}

func CreateShader(vertex, fragment string) (*Shader, error) {
  p, err := glutil.CreateProgram(GL, vertex, fragment)
  if err != nil {
    return nil, err
  }

  //todo? atriblocations https://github.com/pixijs/pixi-gl-core/blob/master/src/shader/compileProgram.js

  s := &Shader{}
  s.Program = p
  s.Attributes = CreateShaderAttributes(s.Program)
  s.UniformData = CreateShaderUniforms(s.Program)
  log.Printf("%+v", s.Attributes)

  return s, nil
}

func (s *Shader) Bind() {
  GL.UseProgram(s.Program)
}

func (s *Shader) Destroy() {
  //todo
}


type ShaderAttrib struct {
  Type string
  Size int
  Location int
}

func CreateShaderAttributes(program *gl.Program) map[string]*ShaderAttrib {
  attributes := make(map[string]*ShaderAttrib)
  total := GL.GetProgramParameteri(program, GL.ACTIVE_ATTRIBUTES)

  for i := 0; i < total; i++ {
    name, size, typ := GL.GetActiveAttrib(program, i)
    log.Println(name, size, typ)
    attributes[name] = &ShaderAttrib{
      mapType(typ),
      size,
      GL.GetAttribLocation(program, name),
    }

  }

  return attributes
}

func (sa *ShaderAttrib) Pointer(typ int, normalized bool, stride, start int) {
  GL.VertexAttribPointer(sa.Location, sa.Size, typ, normalized, stride, 0)
}

func (sa *ShaderAttrib) DefaultPointer() {
  GL.VertexAttribPointer(sa.Location, sa.Size, GL.FLOAT, false, 0, 0)
}

type ShaderUniform struct {
  Type string
  Size int
  Location int
  Value interface{}
}

func CreateShaderUniforms(program *gl.Program) map[string]*ShaderUniform {
  uniforms := make(map[string]*ShaderUniform)

  return uniforms
}

func mapType(typ int) string {
  types := map[int]string{
    GL.FLOAT: "float",
    GL.FLOAT_VEC2: "vec2",
    GL.FLOAT_VEC3: "vec3",
    GL.FLOAT_VEC4: "vec4",

    GL.INT: "int",
    GL.INT_VEC2: "ivec2",
    GL.INT_VEC3: "ivec3",
    GL.INT_VEC4: "ivec4",

    GL.BOOL: "bool",
    GL.BOOL_VEC2: "bvec2",
    GL.BOOL_VEC3: "bvec3",
    GL.BOOL_VEC4: "bvec4",

    GL.FLOAT_MAT2: "mat2",
    GL.FLOAT_MAT3: "mat3",
    GL.FLOAT_MAT4: "mat4",

    GL.SAMPLER_2D: "sampler2D",
  }

  v, ok := types[typ]
  if !ok {
    return ""
  }

  return v
}

var textureVert = `
#ifdef GL_ES
precision highp float;
#endif
attribute vec2 a_vertexPosition;
attribute vec2 a_textureCoord;
attribute vec4 a_color;
attribute float a_textureId;

uniform mat3 projectionMatrix;

varying vec2 v_textureCoord;
varying vec4 v_color;
varying float v_textureId;

void main(void){
    gl_Position = vec4((projectionMatrix * vec3(a_vertexPosition, 1.0)).xy, 0.0, 1.0);

    v_textureCoord = a_textureCoord;
    v_textureId = a_textureId;
    v_color = vec4(a_color.rgb * a_color.a, a_color.a);
}
`

var multiTextureFrag = `
#ifdef GL_ES
precision highp float;
#endif
varying vec2 v_textureCoord;
varying vec4 v_color;
varying float v_textureId;
uniform sampler2D u_samplers[%num%];

void main(){
	vec4 color;
	float textureId = floor(v_textureId+0.5);
	%loop%
	gl_FragColor = color * v_color;
}
`

func CreateMultiTextureShader() *gl.Program {
  frag := strings.Replace(multiTextureFrag, "%num%", strconv.Itoa(MAX_TEXTURES), -1)
  frag = strings.Replace(frag, "%loop%", generateSamplers(), -1)

  p, err := glutil.CreateProgram(GL, textureVert, frag)
  if err != nil {
    log.Fatal(err)
  }

  //todo https://github.com/pixijs/pixi.js/blob/dev/src/core/sprites/webgl/generateMultiTextureShader.js
  return p
}

func generateSamplers() string {
  str := "\n\n"

  for i := 0; i < MAX_TEXTURES; i++ {
    n := strconv.Itoa(i)

    if i > 0 {
      str += "\nelse "
    }

    if i < MAX_TEXTURES -1 {
      str += "if(textureId == " + n + ".0)"
    }

    str += `{
      color = texture2D(u_samplers[` + n + `], v_textureCoord);
    }
    `
  }

  str += "\n\n"

  return str
}
