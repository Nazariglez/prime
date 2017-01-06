// Created by nazarigonzalez on 6/1/17.

// +build !js

package loop

import (
  "time"
)

func (l *loopContext) update() {
  var now, delta int64
  for _ = range l.ticker.C {
    now = time.Now().UnixNano()

    mu.Lock()

    l.time += now - l.last
    delta = l.time - l.lastTime

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

    mu.Unlock()

    l.tickFn(l.lastDelta)
  }
}
