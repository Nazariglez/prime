/**
 * Created by nazarigonzalez on 29/12/16.
 */

package prime

type PrimeOptions struct {
  Title string
  Width, Height int
  Background Vec4
}

var defaultOptions = &PrimeOptions{
  Title: "A Shiny Prime Game!",
  Width: 600,
  Height: 1024,
  Background: Vec4{0.9, 0.1, 0.4, 0},
}

func parseOptions(o *PrimeOptions) *PrimeOptions {
  if o == nil {
    return defaultOptions
  } else {
    //todo parse options with default options
    return defaultOptions
  }
}

type Vec4 [4]float32