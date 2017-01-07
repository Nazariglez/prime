// The original code was created by Joseph Hager under BSD License. (https://github.com/ajhager)
// https://github.com/ajhager/webgl/blob/master/webgl_gl2.go

// +build !js
// +build !android

package gl

import (
	"image"
	"reflect"

	gl2 "github.com/go-gl/gl/v2.1/gl"
	"strings"
)

type Texture struct{ uint32 }
type Buffer struct{ uint32 }
type FrameBuffer struct{ uint32 }
type RenderBuffer struct{ uint32 }
type Program struct{ uint32 }
type UniformLocation struct{ int32 }
type Shader struct{ uint32 }

type Context struct {
	BaseContext
}

func NewContext() (*Context, error) {
	if err := gl2.Init(); err != nil {
		return nil, err
	}
	return &Context{
		BaseContext{
			ARRAY_BUFFER:               gl2.ARRAY_BUFFER,
			ARRAY_BUFFER_BINDING:       gl2.ARRAY_BUFFER_BINDING,
			ATTACHED_SHADERS:           gl2.ATTACHED_SHADERS,
			BACK:                       gl2.BACK,
			BLEND:                      gl2.BLEND,
			BLEND_COLOR:                gl2.BLEND_COLOR,
			BLEND_DST_ALPHA:            gl2.BLEND_DST_ALPHA,
			BLEND_DST_RGB:              gl2.BLEND_DST_RGB,
			BLEND_EQUATION:             gl2.BLEND_EQUATION,
			BLEND_EQUATION_ALPHA:       gl2.BLEND_EQUATION_ALPHA,
			BLEND_EQUATION_RGB:         gl2.BLEND_EQUATION_RGB,
			BLEND_SRC_ALPHA:            gl2.BLEND_SRC_ALPHA,
			BLEND_SRC_RGB:              gl2.BLEND_SRC_RGB,
			BLUE_BITS:                  gl2.BLUE_BITS,
			BOOL:                       gl2.BOOL,
			BOOL_VEC2:                  gl2.BOOL_VEC2,
			BOOL_VEC3:                  gl2.BOOL_VEC3,
			BOOL_VEC4:                  gl2.BOOL_VEC4,
			BUFFER_SIZE:                gl2.BUFFER_SIZE,
			BUFFER_USAGE:               gl2.BUFFER_USAGE,
			BYTE:                       gl2.BYTE,
			CCW:                        gl2.CCW,
			CLAMP_TO_EDGE:              gl2.CLAMP_TO_EDGE,
			COLOR_ATTACHMENT0:          gl2.COLOR_ATTACHMENT0,
			COLOR_BUFFER_BIT:           gl2.COLOR_BUFFER_BIT,
			COLOR_CLEAR_VALUE:          gl2.COLOR_CLEAR_VALUE,
			COLOR_WRITEMASK:            gl2.COLOR_WRITEMASK,
			COMPILE_STATUS:             gl2.COMPILE_STATUS,
			COMPRESSED_TEXTURE_FORMATS: gl2.COMPRESSED_TEXTURE_FORMATS,
			CONSTANT_ALPHA:             gl2.CONSTANT_ALPHA,
			CONSTANT_COLOR:             gl2.CONSTANT_COLOR,
			CULL_FACE:                  gl2.CULL_FACE,
			CULL_FACE_MODE:             gl2.CULL_FACE_MODE,
			CURRENT_PROGRAM:            gl2.CURRENT_PROGRAM,
			CURRENT_VERTEX_ATTRIB:      gl2.CURRENT_VERTEX_ATTRIB,
			CW:                           gl2.CW,
			DECR:                         gl2.DECR,
			DECR_WRAP:                    gl2.DECR_WRAP,
			DELETE_STATUS:                gl2.DELETE_STATUS,
			DEPTH_ATTACHMENT:             gl2.DEPTH_ATTACHMENT,
			DEPTH_BITS:                   gl2.DEPTH_BITS,
			DEPTH_BUFFER_BIT:             gl2.DEPTH_BUFFER_BIT,
			DEPTH_CLEAR_VALUE:            gl2.DEPTH_CLEAR_VALUE,
			DEPTH_COMPONENT:              gl2.DEPTH_COMPONENT,
			DEPTH_COMPONENT16:            gl2.DEPTH_COMPONENT16,
			DEPTH_FUNC:                   gl2.DEPTH_FUNC,
			DEPTH_RANGE:                  gl2.DEPTH_RANGE,
			DEPTH_STENCIL:                gl2.DEPTH_STENCIL,
			DEPTH_STENCIL_ATTACHMENT:     gl2.DEPTH_STENCIL_ATTACHMENT,
			DEPTH_TEST:                   gl2.DEPTH_TEST,
			DEPTH_WRITEMASK:              gl2.DEPTH_WRITEMASK,
			DITHER:                       gl2.DITHER,
			DONT_CARE:                    gl2.DONT_CARE,
			DST_ALPHA:                    gl2.DST_ALPHA,
			DST_COLOR:                    gl2.DST_COLOR,
			DYNAMIC_DRAW:                 gl2.DYNAMIC_DRAW,
			ELEMENT_ARRAY_BUFFER:         gl2.ELEMENT_ARRAY_BUFFER,
			ELEMENT_ARRAY_BUFFER_BINDING: gl2.ELEMENT_ARRAY_BUFFER_BINDING,
			EQUAL:                                        gl2.EQUAL,
			FASTEST:                                      gl2.FASTEST,
			FLOAT:                                        gl2.FLOAT,
			FLOAT_MAT2:                                   gl2.FLOAT_MAT2,
			FLOAT_MAT3:                                   gl2.FLOAT_MAT3,
			FLOAT_MAT4:                                   gl2.FLOAT_MAT4,
			FLOAT_VEC2:                                   gl2.FLOAT_VEC2,
			FLOAT_VEC3:                                   gl2.FLOAT_VEC3,
			FLOAT_VEC4:                                   gl2.FLOAT_VEC4,
			FRAGMENT_SHADER:                              gl2.FRAGMENT_SHADER,
			FRAMEBUFFER:                                  gl2.FRAMEBUFFER,
			FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:           gl2.FRAMEBUFFER_ATTACHMENT_OBJECT_NAME,
			FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:           gl2.FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE,
			FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE: gl2.FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE,
			FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:         gl2.FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL,
			FRAMEBUFFER_BINDING:                          gl2.FRAMEBUFFER_BINDING,
			FRAMEBUFFER_COMPLETE:                         gl2.FRAMEBUFFER_COMPLETE,
			FRAMEBUFFER_INCOMPLETE_ATTACHMENT:            gl2.FRAMEBUFFER_INCOMPLETE_ATTACHMENT,
			FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:    gl2.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT,
			FRAMEBUFFER_UNSUPPORTED:                      gl2.FRAMEBUFFER_UNSUPPORTED,
			FRONT:                         gl2.FRONT,
			FRONT_AND_BACK:                gl2.FRONT_AND_BACK,
			FRONT_FACE:                    gl2.FRONT_FACE,
			FUNC_ADD:                      gl2.FUNC_ADD,
			FUNC_REVERSE_SUBTRACT:         gl2.FUNC_REVERSE_SUBTRACT,
			FUNC_SUBTRACT:                 gl2.FUNC_SUBTRACT,
			GENERATE_MIPMAP_HINT:          gl2.GENERATE_MIPMAP_HINT,
			GEQUAL:                        gl2.GEQUAL,
			GREATER:                       gl2.GREATER,
			GREEN_BITS:                    gl2.GREEN_BITS,
			HIGH_FLOAT:                    gl2.HIGH_FLOAT,
			HIGH_INT:                      gl2.HIGH_INT,
			INCR:                          gl2.INCR,
			INCR_WRAP:                     gl2.INCR_WRAP,
			INFO_LOG_LENGTH:               gl2.INFO_LOG_LENGTH,
			INT:                           gl2.INT,
			INT_VEC2:                      gl2.INT_VEC2,
			INT_VEC3:                      gl2.INT_VEC3,
			INT_VEC4:                      gl2.INT_VEC4,
			INVALID_ENUM:                  gl2.INVALID_ENUM,
			INVALID_FRAMEBUFFER_OPERATION: gl2.INVALID_FRAMEBUFFER_OPERATION,
			INVALID_OPERATION:             gl2.INVALID_OPERATION,
			INVALID_VALUE:                 gl2.INVALID_VALUE,
			INVERT:                        gl2.INVERT,
			KEEP:                          gl2.KEEP,
			LEQUAL:                        gl2.LEQUAL,
			LESS:                          gl2.LESS,
			LINEAR:                        gl2.LINEAR,
			LINEAR_MIPMAP_LINEAR:          gl2.LINEAR_MIPMAP_LINEAR,
			LINEAR_MIPMAP_NEAREST:         gl2.LINEAR_MIPMAP_NEAREST,
			LINES:                            gl2.LINES,
			LINE_LOOP:                        gl2.LINE_LOOP,
			LINE_STRIP:                       gl2.LINE_STRIP,
			LINE_WIDTH:                       gl2.LINE_WIDTH,
			LINK_STATUS:                      gl2.LINK_STATUS,
			LOW_FLOAT:                        gl2.LOW_FLOAT,
			LOW_INT:                          gl2.LOW_INT,
			LUMINANCE:                        gl2.LUMINANCE,
			LUMINANCE_ALPHA:                  gl2.LUMINANCE_ALPHA,
			MAX_COMBINED_TEXTURE_IMAGE_UNITS: gl2.MAX_COMBINED_TEXTURE_IMAGE_UNITS,
			MAX_CUBE_MAP_TEXTURE_SIZE:        gl2.MAX_CUBE_MAP_TEXTURE_SIZE,
			MAX_FRAGMENT_UNIFORM_VECTORS:     gl2.MAX_FRAGMENT_UNIFORM_VECTORS,
			MAX_RENDERBUFFER_SIZE:            gl2.MAX_RENDERBUFFER_SIZE,
			MAX_TEXTURE_IMAGE_UNITS:          gl2.MAX_TEXTURE_IMAGE_UNITS,
			MAX_TEXTURE_SIZE:                 gl2.MAX_TEXTURE_SIZE,
			MAX_VARYING_VECTORS:              gl2.MAX_VARYING_VECTORS,
			MAX_VERTEX_ATTRIBS:               gl2.MAX_VERTEX_ATTRIBS,
			MAX_VERTEX_TEXTURE_IMAGE_UNITS:   gl2.MAX_VERTEX_TEXTURE_IMAGE_UNITS,
			MAX_VERTEX_UNIFORM_VECTORS:       gl2.MAX_VERTEX_UNIFORM_VECTORS,
			MAX_VIEWPORT_DIMS:                gl2.MAX_VIEWPORT_DIMS,
			MEDIUM_FLOAT:                     gl2.MEDIUM_FLOAT,
			MEDIUM_INT:                       gl2.MEDIUM_INT,
			MIRRORED_REPEAT:                  gl2.MIRRORED_REPEAT,
			NEAREST:                          gl2.NEAREST,
			NEAREST_MIPMAP_LINEAR:            gl2.NEAREST_MIPMAP_LINEAR,
			NEAREST_MIPMAP_NEAREST:           gl2.NEAREST_MIPMAP_NEAREST,
			NEVER:                          gl2.NEVER,
			NICEST:                         gl2.NICEST,
			NONE:                           gl2.NONE,
			NOTEQUAL:                       gl2.NOTEQUAL,
			NO_ERROR:                       gl2.NO_ERROR,
			NUM_COMPRESSED_TEXTURE_FORMATS: gl2.NUM_COMPRESSED_TEXTURE_FORMATS,
			ONE: gl2.ONE,
			ONE_MINUS_CONSTANT_ALPHA:     gl2.ONE_MINUS_CONSTANT_ALPHA,
			ONE_MINUS_CONSTANT_COLOR:     gl2.ONE_MINUS_CONSTANT_COLOR,
			ONE_MINUS_DST_ALPHA:          gl2.ONE_MINUS_DST_ALPHA,
			ONE_MINUS_DST_COLOR:          gl2.ONE_MINUS_DST_COLOR,
			ONE_MINUS_SRC_ALPHA:          gl2.ONE_MINUS_SRC_ALPHA,
			ONE_MINUS_SRC_COLOR:          gl2.ONE_MINUS_SRC_COLOR,
			OUT_OF_MEMORY:                gl2.OUT_OF_MEMORY,
			PACK_ALIGNMENT:               gl2.PACK_ALIGNMENT,
			POINTS:                       gl2.POINTS,
			POLYGON_OFFSET_FACTOR:        gl2.POLYGON_OFFSET_FACTOR,
			POLYGON_OFFSET_FILL:          gl2.POLYGON_OFFSET_FILL,
			POLYGON_OFFSET_UNITS:         gl2.POLYGON_OFFSET_UNITS,
			RED_BITS:                     gl2.RED_BITS,
			RENDERBUFFER:                 gl2.RENDERBUFFER,
			RENDERBUFFER_ALPHA_SIZE:      gl2.RENDERBUFFER_ALPHA_SIZE,
			RENDERBUFFER_BINDING:         gl2.RENDERBUFFER_BINDING,
			RENDERBUFFER_BLUE_SIZE:       gl2.RENDERBUFFER_BLUE_SIZE,
			RENDERBUFFER_DEPTH_SIZE:      gl2.RENDERBUFFER_DEPTH_SIZE,
			RENDERBUFFER_GREEN_SIZE:      gl2.RENDERBUFFER_GREEN_SIZE,
			RENDERBUFFER_HEIGHT:          gl2.RENDERBUFFER_HEIGHT,
			RENDERBUFFER_INTERNAL_FORMAT: gl2.RENDERBUFFER_INTERNAL_FORMAT,
			RENDERBUFFER_RED_SIZE:        gl2.RENDERBUFFER_RED_SIZE,
			RENDERBUFFER_STENCIL_SIZE:    gl2.RENDERBUFFER_STENCIL_SIZE,
			RENDERBUFFER_WIDTH:           gl2.RENDERBUFFER_WIDTH,
			RENDERER:                     gl2.RENDERER,
			REPEAT:                       gl2.REPEAT,
			REPLACE:                      gl2.REPLACE,
			RGB:                          gl2.RGB,
			RGB5_A1:                      gl2.RGB5_A1,
			RGB565:                       gl2.RGB565,
			RGBA:                         gl2.RGBA,
			RGBA4:                        gl2.RGBA4,
			SAMPLER_2D:                   gl2.SAMPLER_2D,
			SAMPLER_CUBE:                 gl2.SAMPLER_CUBE,
			SAMPLES:                      gl2.SAMPLES,
			SAMPLE_ALPHA_TO_COVERAGE:     gl2.SAMPLE_ALPHA_TO_COVERAGE,
			SAMPLE_BUFFERS:               gl2.SAMPLE_BUFFERS,
			SAMPLE_COVERAGE:              gl2.SAMPLE_COVERAGE,
			SAMPLE_COVERAGE_INVERT:       gl2.SAMPLE_COVERAGE_INVERT,
			SAMPLE_COVERAGE_VALUE:        gl2.SAMPLE_COVERAGE_VALUE,
			SCISSOR_BOX:                  gl2.SCISSOR_BOX,
			SCISSOR_TEST:                 gl2.SCISSOR_TEST,
			SHADER_COMPILER:              gl2.SHADER_COMPILER,
			SHADER_SOURCE_LENGTH:         gl2.SHADER_SOURCE_LENGTH,
			SHADER_TYPE:                  gl2.SHADER_TYPE,
			SHADING_LANGUAGE_VERSION:     gl2.SHADING_LANGUAGE_VERSION,
			SHORT:                        gl2.SHORT,
			SRC_ALPHA:                    gl2.SRC_ALPHA,
			SRC_ALPHA_SATURATE:           gl2.SRC_ALPHA_SATURATE,
			SRC_COLOR:                    gl2.SRC_COLOR,
			STATIC_DRAW:                  gl2.STATIC_DRAW,
			STENCIL_ATTACHMENT:           gl2.STENCIL_ATTACHMENT,
			STENCIL_BACK_FAIL:            gl2.STENCIL_BACK_FAIL,
			STENCIL_BACK_FUNC:            gl2.STENCIL_BACK_FUNC,
			STENCIL_BACK_PASS_DEPTH_FAIL: gl2.STENCIL_BACK_PASS_DEPTH_FAIL,
			STENCIL_BACK_PASS_DEPTH_PASS: gl2.STENCIL_BACK_PASS_DEPTH_PASS,
			STENCIL_BACK_REF:             gl2.STENCIL_BACK_REF,
			STENCIL_BACK_VALUE_MASK:      gl2.STENCIL_BACK_VALUE_MASK,
			STENCIL_BACK_WRITEMASK:       gl2.STENCIL_BACK_WRITEMASK,
			STENCIL_BITS:                 gl2.STENCIL_BITS,
			STENCIL_BUFFER_BIT:           gl2.STENCIL_BUFFER_BIT,
			STENCIL_CLEAR_VALUE:          gl2.STENCIL_CLEAR_VALUE,
			STENCIL_FAIL:                 gl2.STENCIL_FAIL,
			STENCIL_FUNC:                 gl2.STENCIL_FUNC,
			STENCIL_INDEX:                gl2.STENCIL_INDEX,
			STENCIL_INDEX8:               gl2.STENCIL_INDEX8,
			STENCIL_PASS_DEPTH_FAIL:      gl2.STENCIL_PASS_DEPTH_FAIL,
			STENCIL_PASS_DEPTH_PASS:      gl2.STENCIL_PASS_DEPTH_PASS,
			STENCIL_REF:                  gl2.STENCIL_REF,
			STENCIL_TEST:                 gl2.STENCIL_TEST,
			STENCIL_VALUE_MASK:           gl2.STENCIL_VALUE_MASK,
			STENCIL_WRITEMASK:            gl2.STENCIL_WRITEMASK,
			STREAM_DRAW:                  gl2.STREAM_DRAW,
			SUBPIXEL_BITS:                gl2.SUBPIXEL_BITS,
			TEXTURE:                      gl2.TEXTURE,
			TEXTURE0:                     gl2.TEXTURE0,
			TEXTURE1:                     gl2.TEXTURE1,
			TEXTURE2:                     gl2.TEXTURE2,
			TEXTURE3:                     gl2.TEXTURE3,
			TEXTURE4:                     gl2.TEXTURE4,
			TEXTURE5:                     gl2.TEXTURE5,
			TEXTURE6:                     gl2.TEXTURE6,
			TEXTURE7:                     gl2.TEXTURE7,
			TEXTURE8:                     gl2.TEXTURE8,
			TEXTURE9:                     gl2.TEXTURE9,
			TEXTURE10:                    gl2.TEXTURE10,
			TEXTURE11:                    gl2.TEXTURE11,
			TEXTURE12:                    gl2.TEXTURE12,
			TEXTURE13:                    gl2.TEXTURE13,
			TEXTURE14:                    gl2.TEXTURE14,
			TEXTURE15:                    gl2.TEXTURE15,
			TEXTURE16:                    gl2.TEXTURE16,
			TEXTURE17:                    gl2.TEXTURE17,
			TEXTURE18:                    gl2.TEXTURE18,
			TEXTURE19:                    gl2.TEXTURE19,
			TEXTURE20:                    gl2.TEXTURE20,
			TEXTURE21:                    gl2.TEXTURE21,
			TEXTURE22:                    gl2.TEXTURE22,
			TEXTURE23:                    gl2.TEXTURE23,
			TEXTURE24:                    gl2.TEXTURE24,
			TEXTURE25:                    gl2.TEXTURE25,
			TEXTURE26:                    gl2.TEXTURE26,
			TEXTURE27:                    gl2.TEXTURE27,
			TEXTURE28:                    gl2.TEXTURE28,
			TEXTURE29:                    gl2.TEXTURE29,
			TEXTURE30:                    gl2.TEXTURE30,
			TEXTURE31:                    gl2.TEXTURE31,
			TEXTURE_2D:                   gl2.TEXTURE_2D,
			TEXTURE_BINDING_2D:           gl2.TEXTURE_BINDING_2D,
			TEXTURE_BINDING_CUBE_MAP:     gl2.TEXTURE_BINDING_CUBE_MAP,
			TEXTURE_CUBE_MAP:             gl2.TEXTURE_CUBE_MAP,
			TEXTURE_CUBE_MAP_NEGATIVE_X:  gl2.TEXTURE_CUBE_MAP_NEGATIVE_X,
			TEXTURE_CUBE_MAP_NEGATIVE_Y:  gl2.TEXTURE_CUBE_MAP_NEGATIVE_Y,
			TEXTURE_CUBE_MAP_NEGATIVE_Z:  gl2.TEXTURE_CUBE_MAP_NEGATIVE_Z,
			TEXTURE_CUBE_MAP_POSITIVE_X:  gl2.TEXTURE_CUBE_MAP_POSITIVE_X,
			TEXTURE_CUBE_MAP_POSITIVE_Y:  gl2.TEXTURE_CUBE_MAP_POSITIVE_Y,
			TEXTURE_CUBE_MAP_POSITIVE_Z:  gl2.TEXTURE_CUBE_MAP_POSITIVE_Z,
			TEXTURE_MAG_FILTER:           gl2.TEXTURE_MAG_FILTER,
			TEXTURE_MIN_FILTER:           gl2.TEXTURE_MIN_FILTER,
			TEXTURE_WRAP_S:               gl2.TEXTURE_WRAP_S,
			TEXTURE_WRAP_T:               gl2.TEXTURE_WRAP_T,
			TRIANGLES:                    gl2.TRIANGLES,
			TRIANGLE_FAN:                 gl2.TRIANGLE_FAN,
			TRIANGLE_STRIP:               gl2.TRIANGLE_STRIP,
			UNPACK_ALIGNMENT:             gl2.UNPACK_ALIGNMENT,
			UNSIGNED_BYTE:                gl2.UNSIGNED_BYTE,
			UNSIGNED_INT:                 gl2.UNSIGNED_INT,
			UNSIGNED_SHORT:               gl2.UNSIGNED_SHORT,
			UNSIGNED_SHORT_4_4_4_4:       gl2.UNSIGNED_SHORT_4_4_4_4,
			UNSIGNED_SHORT_5_5_5_1:       gl2.UNSIGNED_SHORT_5_5_5_1,
			UNSIGNED_SHORT_5_6_5:         gl2.UNSIGNED_SHORT_5_6_5,
			VALIDATE_STATUS:              gl2.VALIDATE_STATUS,
			VENDOR:                       gl2.VENDOR,
			VERSION:                      gl2.VERSION,
			VERTEX_ATTRIB_ARRAY_BUFFER_BINDING: gl2.VERTEX_ATTRIB_ARRAY_BUFFER_BINDING,
			VERTEX_ATTRIB_ARRAY_ENABLED:        gl2.VERTEX_ATTRIB_ARRAY_ENABLED,
			VERTEX_ATTRIB_ARRAY_NORMALIZED:     gl2.VERTEX_ATTRIB_ARRAY_NORMALIZED,
			VERTEX_ATTRIB_ARRAY_POINTER:        gl2.VERTEX_ATTRIB_ARRAY_POINTER,
			VERTEX_ATTRIB_ARRAY_SIZE:           gl2.VERTEX_ATTRIB_ARRAY_SIZE,
			VERTEX_ATTRIB_ARRAY_STRIDE:         gl2.VERTEX_ATTRIB_ARRAY_STRIDE,
			VERTEX_ATTRIB_ARRAY_TYPE:           gl2.VERTEX_ATTRIB_ARRAY_TYPE,
			VERTEX_SHADER:                      gl2.VERTEX_SHADER,
			VIEWPORT:                           gl2.VIEWPORT,
			ZERO:                               gl2.ZERO,
		},
	}, nil
}

func (c *Context) CreateShader(typ int) *Shader {
	shader := &Shader{gl2.CreateShader(uint32(typ))}
	return shader
}

func (c *Context) ShaderSource(shader *Shader, source string) {
	glsource, free := gl2.Strs(source + "\x00")
	gl2.ShaderSource(shader.uint32, 1, glsource, nil)
	free()
}

func (c *Context) CompileShader(shader *Shader) {
	gl2.CompileShader(shader.uint32)
}

func (c *Context) DeleteShader(shader *Shader) {
	gl2.DeleteShader(shader.uint32)
}

func (c *Context) CreateProgram() *Program {
	return &Program{gl2.CreateProgram()}
}

func (c *Context) DeleteProgram(program *Program) {
	gl2.DeleteProgram(program.uint32)
}

func (c *Context) AttachShader(program *Program, shader *Shader) {
	gl2.AttachShader(program.uint32, shader.uint32)
}

func (c *Context) GetShaderParameterb(shader *Shader, pname int) bool {
	var r int32
	gl2.GetShaderiv(shader.uint32, uint32(pname), &r)
	return r == gl2.TRUE
}

func (c *Context) GetProgramParameterb(program *Program, pname int) bool {
	var r int32
	gl2.GetProgramiv(program.uint32, uint32(pname), &r)
	return r == gl2.TRUE
}

func (c *Context) GetShaderInfoLog(shader *Shader) string {
	var l int32
	gl2.GetShaderiv(shader.uint32, gl2.INFO_LOG_LENGTH, &l)

	s := strings.Repeat("\x00", int(l+1))
	gl2.GetShaderInfoLog(shader.uint32, l, nil, gl2.Str(s))
	return s
}

func (c *Context) GetProgramInfoLog(program *Program) string {
	var l int32
	gl2.GetProgramiv(program.uint32, gl2.INFO_LOG_LENGTH, &l)

	s := strings.Repeat("\x00", int(l+1))
	gl2.GetProgramInfoLog(program.uint32, l, nil, gl2.Str(s))
	return s
}

func (c *Context) LinkProgram(program *Program) {
	gl2.LinkProgram(program.uint32)
}

func (c *Context) CreateTexture() *Texture {
	var loc uint32
	gl2.GenTextures(1, &loc)
	return &Texture{loc}
}

func (c *Context) DeleteTexture(texture *Texture) {
	gl2.DeleteTextures(1, &[]uint32{texture.uint32}[0])
}

func (c *Context) BindTexture(target int, texture *Texture) {
	if texture == nil {
		gl2.BindTexture(uint32(target), 0)
		return
	}
	gl2.BindTexture(uint32(target), texture.uint32)
}

func (c *Context) TexParameteri(target int, pname int, param int) {
	gl2.TexParameteri(uint32(target), uint32(pname), int32(param))
}

func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, data interface{}) {
	var pix []uint8
	width := 0
	height := 0
	if data == nil {
		pix = nil
	} else {
		img := data.(*image.NRGBA)
		width = img.Bounds().Dx()
		height = img.Bounds().Dy()
		pix = img.Pix
	}
	gl2.TexImage2D(uint32(target), int32(level), int32(internalFormat), int32(width), int32(height), int32(0), uint32(format), uint32(kind), gl2.Ptr(pix))
}

func (c *Context) GetAttribLocation(program *Program, name string) int {
	return int(gl2.GetAttribLocation(program.uint32, gl2.Str(name+"\x00")))
}

func (c *Context) GetUniformLocation(program *Program, name string) *UniformLocation {
	return &UniformLocation{gl2.GetUniformLocation(program.uint32, gl2.Str(name+"\x00"))}
}

func (c *Context) GetError() int {
	return int(gl2.GetError())
}

func (c *Context) CreateBuffer() *Buffer {
	var loc uint32
	gl2.GenBuffers(1, &loc)
	return &Buffer{loc}
}

func (c *Context) BindBuffer(target int, buffer *Buffer) {
	if buffer == nil {
		gl2.BindBuffer(uint32(target), 0)
		return
	}
	gl2.BindBuffer(uint32(target), buffer.uint32)
}

func (c *Context) BufferData(target int, data interface{}, usage int) {
	s := uintptr(reflect.ValueOf(data).Len()) * reflect.TypeOf(data).Elem().Size()
	gl2.BufferData(uint32(target), int(s), gl2.Ptr(data), uint32(usage))
}

func (c *Context) EnableVertexAttribArray(index int) {
	gl2.EnableVertexAttribArray(uint32(index))
}

func (c *Context) DisableVertexAttribArray(index int) {
	gl2.DisableVertexAttribArray(uint32(index))
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	gl2.VertexAttribPointer(uint32(index), int32(size), uint32(typ), normal, int32(stride), gl2.PtrOffset(offset))
}

func (c *Context) Enable(flag int) {
	gl2.Enable(uint32(flag))
}

func (c *Context) Disable(flag int) {
	gl2.Disable(uint32(flag))
}

func (c *Context) BlendFunc(src, dst int) {
	gl2.BlendFunc(uint32(src), uint32(dst))
}

func (c *Context) UniformMatrix4fv(location *UniformLocation, transpose bool, value []float32) {
	// TODO: count value of 1 is currently hardcoded.
	//       Perhaps it should be len(value) / 16 or something else?
	//       In OpenGL 2.1 it is a manually supplied parameter, but WebGL does not have it.
	//       Not sure if WebGL automatically deduces it and supports count values greater than 1, or if 1 is always assumed.
	gl2.UniformMatrix4fv(location.int32, 1, transpose, &value[0])
}

func (c *Context) UseProgram(program *Program) {
	if program == nil {
		gl2.UseProgram(0)
		return
	}
	gl2.UseProgram(program.uint32)
}

func (c *Context) ValidateProgram(program *Program) {
	if program == nil {
		gl2.ValidateProgram(0)
		return
	}
	gl2.ValidateProgram(program.uint32)
}

func (c *Context) Uniform1i(location *UniformLocation, x int) {
	gl2.Uniform1i(location.int32, int32(x))
}

func (c *Context) Uniform1f(location *UniformLocation, x float32) {
	gl2.Uniform1f(location.int32, x)
}

func (c *Context) Uniform2f(location *UniformLocation, x, y float32) {
	gl2.Uniform2f(location.int32, x, y)
}

func (c *Context) BufferSubData(target int, offset int, data interface{}) {
	size := uintptr(reflect.ValueOf(data).Len()) * reflect.TypeOf(data).Elem().Size()
	gl2.BufferSubData(uint32(target), offset, int(size), gl2.Ptr(data))
}

func (c *Context) DrawArrays(mode, first, count int) {
	gl2.DrawArrays(uint32(mode), int32(first), int32(count))
}

func (c *Context) DrawElements(mode, count, typ, offset int) {
	gl2.DrawElements(uint32(mode), int32(count), uint32(typ), gl2.PtrOffset(offset))
}

func (c *Context) ClearColor(r, g, b, a float32) {
	gl2.ClearColor(r, g, b, a)
}

func (c *Context) Viewport(x, y, width, height int) {
	gl2.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func (c *Context) Clear(flags int) {
	gl2.Clear(uint32(flags))
}
