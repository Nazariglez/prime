// Created by nazarigonzalez on 29/12/16.

package prime

type PrimeOptions struct {
	Title         string
	Width, Height int
	Background    *Vec4
}

var defaultOptions = &PrimeOptions{
	Title:      "A Shiny Prime Game!",
	Width:      600,
	Height:     1024,
	Background: &Vec4{0.3, 0.3, 0.1, 0.0},
}

func parseOptions(o *PrimeOptions) *PrimeOptions {
	if o == nil {
		return defaultOptions
	} else {

		if o.Title == "" { o.Title = defaultOptions.Title }
		if o.Width == 0 { o.Width = defaultOptions.Width }
		if o.Height == 0 { o.Height = defaultOptions.Height }
		if o.Background == nil { o.Background = defaultOptions.Background }

		return o
	}
}

type Vec4 [4]float32
