// Created by nazarigonzalez on 6/1/17.

package loop

import (
  "time"
  "sync"
  "errors"
)

var currentLoop *loopContext
var mu sync.RWMutex

type loopContext struct {
  fps, currentFps, lastDelta float64
  nanoFps time.Duration
  isRunning bool
  last, lastTime, time int64
  ticker *time.Ticker
  tickFn func (delta float64)

  fpsIndex int
  fpsTotalTime float64

}

func Run(update func(delta float64)) error {
  if currentLoop != nil {
    return errors.New("Can not call loop.Run twice, the game is already running.")
  }

  currentLoop = &loopContext{
    tickFn: update,
  }

  currentLoop.setFPS(60)
  currentLoop.start()
  return nil
}

func (l *loopContext) start() {
  if l.isRunning {
    return
  }

  l.last = time.Now().UnixNano()
  l.isRunning = true
  l.ticker = time.NewTicker(l.nanoFps)

  go l.update()
}

func (l *loopContext) stop() {
  if !l.isRunning {
    return
  }

  l.ticker.Stop()
  l.isRunning = false
}

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

func (l *loopContext) setFPS(fps float64) {
  restart := false
  if l.isRunning {
    l.stop()
    restart = true
  }

  l.fps = fps
  l.nanoFps = time.Duration((1/fps)*1e9) * time.Nanosecond

  if restart {
    l.start()
  }
}

func Start() {
  defer mu.Unlock()
  mu.Lock()

  if currentLoop != nil {
    currentLoop.start()
  }
}

func Stop() {
  defer mu.Unlock()
  mu.Lock()

  if currentLoop != nil {
    currentLoop.stop()
  }
}

func SetFPS(fps float64) {
  defer mu.Unlock()
  mu.Lock()

  if currentLoop != nil {
    currentLoop.setFPS(fps)
  }
}

func GetFPS() float64 {
  defer mu.RUnlock()
  mu.RLock()

  if currentLoop != nil {
    return currentLoop.fps
  }

  return 0
}

func GetRealFPS() float64 {
  defer mu.RUnlock()
  mu.RLock()

  if currentLoop != nil {
    return currentLoop.currentFps
  }

  return 0
}