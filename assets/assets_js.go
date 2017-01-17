// Created by nazarigonzalez on 16/1/17.

// +build js

package assets

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/xhr"
	"prime/gfx"
)

func LoadImage(imgFile string) (*gfx.Image, error) {
	req := xhr.NewRequest("GET", "assets/" + imgFile) //todo path resolve
	req.ResponseType = "blob"

	if err := req.Send(nil); err != nil {
		return nil, err
	}

	if !(req.ReadyState == 4 && req.Status == 200) {
		return nil, errors.New("Unable to load the image: " + imgFile)
	}

	w := make(chan bool, 1)
	img := js.Global.Get("document").Call("createElement", "img")
	img.Set("onload", func() {
		js.Global.Get("window").Get("URL").Call("revokeObjectURL", img.Get("src"))
		w <- true
	})

	img.Set("src", js.Global.Get("window").Get("URL").Call("createObjectURL", req.Response))

	<-w
	js.Global.Get("document").Get("body").Call("appendChild", img)
	return gfx.NewImage(img), nil
}
