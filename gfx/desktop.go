// +build !js
// +build !android

/**
 * Created by nazarigonzalez on 29/12/16.
 */

package gfx

import (
  "github.com/go-gl/gl/v2.1/gl"
  "github.com/go-gl/glfw/v3.2/glfw"
  "log"
  "strings"
  "fmt"
  "github.com/pkg/errors"
)

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


func initialize() error {
  log.Println("Desktop initialized")

  if err := glfw.Init(); err != nil {
    return err
  }

  defer glfw.Terminate()

  glfw.WindowHint(glfw.Samples, 4)
  glfw.WindowHint(glfw.ContextVersionMajor, 2)
  glfw.WindowHint(glfw.ContextVersionMinor, 1)

  window, err := glfw.CreateWindow(gfxWidth, gfxHeight, gfxTitle, nil, nil)
  if err != nil {
    return err
  }

  window.MakeContextCurrent()
  if err := gl.Init(); err != nil {
    return err
  }

  var vertexArrayId uint32
  gl.GenVertexArrays(1, &vertexArrayId)
  gl.BindVertexArray(vertexArrayId)

  gVertexBufferData := []float32{
    -1, -1, 0,
    1, -1, 0,
    0, 1, 0,
  }

  var vertexBuffer uint32
  gl.GenBuffers(1, &vertexBuffer)
  gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)

  sizeOfData := len(gVertexBufferData)*4
  gl.BufferData(gl.ARRAY_BUFFER, sizeOfData, gl.Ptr(gVertexBufferData), gl.STATIC_DRAW)

  programID, err := CreateProgram(vertexShader, fragmentShader)
  if err != nil {
    return err
  }
  gl.BindAttribLocation(programID, 0, gl.Str("vertexPosition_modelspace\x00"))


  gl.ClearColor(gfxBg[0], gfxBg[1], gfxBg[2], gfxBg[3])

  for !window.ShouldClose() {
    //draw here

    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    gl.UseProgram(programID)

    //1rst attribute buffer : vertices
    gl.EnableVertexAttribArray(0)
    gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
    gl.VertexAttribPointer(
      0, //attribute 0. No particular reason fo 0, but must match the layout in the shader
      3, //size
      gl.FLOAT, //type
      false, //normalized?
      0, //stride
      nil, //array buffer offset
    )

    //draw the triangle!
    gl.DrawArrays(gl.TRIANGLES, 0, 3) //starting from vertex 0; 3 vertices total -> 1 triangle
    gl.DisableVertexAttribArray(0)

    window.SwapBuffers()
    glfw.PollEvents()
  }

  return nil
}


func CreateProgram(v, f string) (uint32, error) {
  var invalid bool

  //create the shaders
  vertexShaderID := gl.CreateShader(gl.VERTEX_SHADER)
  fragmentShaderID := gl.CreateShader(gl.FRAGMENT_SHADER)

  //read the shader code from the file
  vertexShaderCode := v+"\x00"
  fragmentShaderCode := f+"\x00"

  var result int32
  var infoLogLength int32

  //compile vertex shader
  vertexSourcePointer, free := gl.Strs(vertexShaderCode)
  defer free()

  gl.ShaderSource(vertexShaderID, 1, vertexSourcePointer, nil)
  gl.CompileShader(vertexShaderID)

  //check vertex shader
  gl.GetShaderiv(vertexShaderID, gl.COMPILE_STATUS, &result)
  gl.GetShaderiv(vertexShaderID, gl.INFO_LOG_LENGTH, &infoLogLength)
  if result != gl.TRUE && infoLogLength > 0 {
    errorLog := strings.Repeat("\x00", int(infoLogLength+1))
    gl.GetShaderInfoLog(vertexShaderID, infoLogLength, nil, gl.Str(errorLog))
    fmt.Printf("[%s]:\n%s\n", v, errorLog)
    invalid = true
  }

  //compile fragment shader
  fragmentSourcePointer, free := gl.Strs(fragmentShaderCode)
  defer free()

  gl.ShaderSource(fragmentShaderID, 1, fragmentSourcePointer, nil)
  gl.CompileShader(fragmentShaderID)

  //check fragment shader
  gl.GetShaderiv(fragmentShaderID, gl.COMPILE_STATUS, &result)
  gl.GetShaderiv(fragmentShaderID, gl.INFO_LOG_LENGTH, &infoLogLength)
  if result != gl.TRUE && infoLogLength > 0 {
    errorLog := strings.Repeat("\x00", int(infoLogLength+1))
    gl.GetShaderInfoLog(fragmentShaderID, infoLogLength, nil, gl.Str(errorLog))
    fmt.Printf("[%s]:\n%s\n", f, errorLog)
    invalid = true
  }

  //link the program
  programID := gl.CreateProgram()
  gl.AttachShader(programID, vertexShaderID)
  gl.AttachShader(programID, fragmentShaderID)
  gl.LinkProgram(programID)

  //check the program
  gl.GetProgramiv(programID, gl.LINK_STATUS, &result)
  gl.GetProgramiv(programID, gl.INFO_LOG_LENGTH, &infoLogLength)
  if result != gl.TRUE && infoLogLength > 0 {
    errorLog := strings.Repeat("\x00", int(infoLogLength+1))
    gl.GetProgramInfoLog(programID, infoLogLength, nil, gl.Str(errorLog))
    fmt.Printf("[%s]:\n%s\n", "Program", errorLog)
    invalid = true
  }

  gl.DetachShader(programID, vertexShaderID)
  gl.DetachShader(programID, fragmentShaderID)

  gl.DeleteShader(vertexShaderID)
  gl.DeleteShader(fragmentShaderID)

  if invalid {
    return 0, errors.New("Invalid shader.")
  }

  return programID, nil
}