// Created by nazarigonzalez on 29/12/16.

// +build js

package gfx

import (
	"log"
	"runtime"
	"strconv"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"

	"prime/gfx/gl"
)

var htmlContentLoaded bool

func initialize() error {
	runtime.LockOSThread()
	log.Println("Browser initialized")

	doc := dom.GetWindow().Document().(dom.HTMLDocument)
	doc.SetTitle(gfxTitle)

	w := make(chan error)
	onReady(func() {
		defer close(w)
		log.Println("Document Loaded!!")

		var canvas *dom.HTMLCanvasElement
		if doc.GetElementByID("prime-view") == nil {
			canvas = doc.CreateElement("canvas").(*dom.HTMLCanvasElement)
			canvas.Set("id", "prime-view")
			doc.Body().AppendChild(canvas)
		} else {
			canvas = doc.GetElementByID("prime-view").(*dom.HTMLCanvasElement)
		}

		canvas.Width = gfxWidth
		canvas.Height = gfxHeight

		scaleCanvas(canvas, gfxScale)

		if err := run(canvas); err != nil {
			w <- err
			return
		}

		w <- nil
	})

	return <-w
}

func run(canvas *dom.HTMLCanvasElement) error {
	attrs := gl.DefaultAttributes()
	attrs.Alpha = false
	attrs.Depth = false
	attrs.PremultipliedAlpha = false
	attrs.PreserveDrawingBuffer = false
	attrs.Antialias = false

	ctx, err := gl.NewContext(canvas.Object, attrs)
	if err != nil {
		return err
	}

	GL = ctx

	go RunSafeReader()
	go OnStart()

	js.Global.Set("ctx", ctx)               //todo remove this
	ctx.Viewport(0, 0, gfxWidth, gfxHeight) //todo fix this, retina issues (removed from here)

	return nil
}

func postRender() {}

func scaleCanvas(canvas *dom.HTMLCanvasElement, typ int) {
	var scale float32
	win := dom.GetWindow()
	wf32 := float32(gfxWidth)
	hf32 := float32(gfxHeight)
	ww := float32(win.InnerWidth()) / wf32
	hh := float32(win.InnerHeight()) / hf32

	log.Println("Scale", typ, BROWSER_SCALE_FIT, typ == BROWSER_SCALE_FIT)

	switch typ {
	case BROWSER_SCALE_FIT:
		if ww < hh {
			scale = ww
		} else {
			scale = hh
		}

		canvas.Style().SetProperty("width", strconv.Itoa(int(wf32*scale))+"px", "")
		canvas.Style().SetProperty("height", strconv.Itoa(int(hf32*scale))+"px", "")

	case BROWSER_SCALE_ASPECT_FILL:
		if ww > hh {
			scale = ww
		} else {
			scale = hh
		}

		canvas.Style().SetProperty("width", strconv.Itoa(int(wf32*scale))+"px", "")
		canvas.Style().SetProperty("height", strconv.Itoa(int(hf32*scale))+"px", "")

	case BROWSER_SCALE_FILL:
		canvas.Style().SetProperty("width", strconv.Itoa(win.InnerWidth())+"px", "")
		canvas.Style().SetProperty("height", strconv.Itoa(win.InnerHeight())+"px", "")

	}

	//todo onresize event
}

func onReady(cb func()) {
	d := js.Global.Get("document")

	if isReadyStateComplete() {
		htmlContentLoaded = true
		cb()
		return
	}

	if d.Get("addEventListener") != nil {
		d.Call("addEventListener", "DOMContentLoaded", onLoad(cb), false)
		js.Global.Call("addEventListener", "load", onLoad(cb), false)
	} else {
		d.Call("attachEvent", "onreadystatechange", func() {
			if isReadyStateComplete() {
				htmlContentLoaded = true
				cb()
			}
		})
		js.Global.Call("attachEvent", "onload", onLoad(cb))
	}

}

func onLoad(cb func()) func() {
	return func() {
		if htmlContentLoaded {
			return
		}

		htmlContentLoaded = true
		cb()
	}
}

func isReadyStateComplete() bool {
	return js.Global.Get("document").Get("readyState").String() == "complete"
}
