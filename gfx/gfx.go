/**
 * Created by nazarigonzalez on 29/12/16.
 */

package gfx

import (
  "runtime"
)

var (
  gfxWidth int
  gfxHeight int
  gfxTitle string
  gfxBg []float32
)

func Init(width, height int, title string, bg []float32) error {
  runtime.LockOSThread()
  gfxWidth = width
  gfxHeight = height
  gfxTitle = title
  gfxBg = bg

  return initialize()
}



