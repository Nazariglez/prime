// Created by nazarigonzalez on 6/1/17.

// +build js

package loop

import (
  "time"
  "github.com/gopherjs/gopherjs/js"
  "log"
)

func (l *loopContext) update() {
  js.Global.Get("window").Call("requestAnimationFrame", l.update)
  now := time.Now().UnixNano() //todo use perfomance.now() because the delta is not precise with UnixNano on browsers
  l.time += now - l.last
  delta := l.time - l.lastTime

  l.lastTime = l.time
  l.last = now
  l.lastDelta = float64(delta)/1e9

  l.fpsTotalTime += l.lastDelta
  l.fpsIndex++

  if l.fpsIndex == 5 {
    l.currentFps = 1/(l.fpsTotalTime/5)
    l.fpsIndex = 0
    l.fpsTotalTime = 0
  }

  go l.tickFn(l.lastDelta)
  log.Println("Delta", l.lastDelta)
}