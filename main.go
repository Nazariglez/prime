/**
 * Created by nazarigonzalez on 29/12/16.
 */

package prime

import (
  "log"
  "prime/gfx"
)

var engineOptions *PrimeOptions

func Initialize(options *PrimeOptions) error {
  log.Println("Starting prime")
  engineOptions = parseOptions(options)
  //extend options with default options

  if err := gfx.Init(engineOptions.Width, engineOptions.Height, engineOptions.Title, engineOptions.Background[:]); err != nil {
    log.Fatal(err)
  }

  return nil
}