// Created by nazarigonzalez on 25/1/17.

package gfx

import (
  "strings"
  "strconv"
  "log"

  "github.com/nazariglez/prime/gfx/gl"
  "github.com/nazariglez/prime/gfx/gl/glutil"
)

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

var textureFrag = `
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
  frag := strings.Replace(textureFrag, "%num%", strconv.Itoa(MAX_TEXTURES), -1)
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