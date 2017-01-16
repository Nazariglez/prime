// Created by nazarigonzalez on 16/1/17.

// +build js

package assets

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/xhr"
	"prime/gfx"
)

func LoadImage(img string) (*gfx.Image, error) {
	req := xhr.NewRequest("GET", img)
	req.ResponseType = "blob"

	if err := req.Send(nil); err != nil {
		return nil, err
	}

	if !(req.ReadyState == 4 && req.Status == 200) {
		return nil, errors.New("Unable to load the image: " + img)
	}

	i := js.Global.Get("document").Call("createElement", "img")
	i.Set("onload", func() {
		js.Global.Get("window").Get("URL").Call("revokeObjectURL", i.Get("src"))
	})

	i.Set("src", js.Global.Get("window").Get("URL").Call("createObjectURL", req.Response))

	js.Global.Get("document").Get("body").Call("appendChild", i)
	return &gfx.Image{}, nil
}