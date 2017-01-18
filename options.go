// Created by nazarigonzalez on 29/12/16.

package prime

import "github.com/nazariglez/prime/gfx"

const (
	BROWSER_SCALE_NONE        = gfx.BROWSER_SCALE_NONE
	BROWSER_SCALE_FIT         = gfx.BROWSER_SCALE_FIT
	BROWSER_SCALE_FILL        = gfx.BROWSER_SCALE_FILL
	BROWSER_SCALE_ASPECT_FILL = gfx.BROWSER_SCALE_ASPECT_FILL
)

type PrimeOptions struct {
	Title         string
	Width, Height int
	Background    *Vec4
	BrowserScale  int
}

var defaultOptions = &PrimeOptions{
	Title:        "A Shiny Prime Game!",
	Width:        600,
	Height:       1024,
	Background:   &Vec4{0.3, 0.3, 0.1, 0.0},
	BrowserScale: BROWSER_SCALE_FIT,
}

func parseOptions(o *PrimeOptions) *PrimeOptions {
	if o == nil {
		return defaultOptions
	} else {

		if o.Title == "" {
			o.Title = defaultOptions.Title
		}
		if o.Width == 0 {
			o.Width = defaultOptions.Width
		}
		if o.Height == 0 {
			o.Height = defaultOptions.Height
		}
		if o.Background == nil {
			o.Background = defaultOptions.Background
		}
		if o.BrowserScale == 0 {
			o.BrowserScale = defaultOptions.BrowserScale
		}

		return o
	}
}

type Vec4 [4]float32
