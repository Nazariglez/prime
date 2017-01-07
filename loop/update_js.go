// Created by nazarigonzalez on 6/1/17.

// +build js

package loop

import (
	"github.com/gopherjs/gopherjs/js"
)

//todo setFps? right now raf dont allow fps
var raf int

func (l *loopContext) start() {
	if l.isRunning {
		return
	}

	l.last = performanceNow()
	l.isRunning = true

	go l.update()
}

func (l *loopContext) stop() {
	if !l.isRunning {
		return
	}

	js.Global.Get("window").Call("cancelAnimationFrame", raf)
	l.isRunning = false
}

func (l *loopContext) update() {
	raf = js.Global.Get("window").Call("requestAnimationFrame", l.update).Int()
	now := performanceNow()
	l.time += now - l.last
	delta := l.time - l.lastTime

	l.lastTime = l.time
	l.last = now
	l.delta = float64(delta) / 1e07

	l.fpsTotalTime += l.delta
	l.fpsIndex++

	if l.fpsIndex == 5 {
		l.currentFps = 1 / (l.fpsTotalTime / 5)
		l.fpsIndex = 0
		l.fpsTotalTime = 0
	}

	go l.tickFn(l.delta)
}

func performanceNow() int64 {
	return int64(js.Global.Get("performance").Call("now").Float() * 1e04)
}
