// Created by nazarigonzalez on 29/12/16.

package prime

import (
	"log"
)

func Init(options *PrimeOptions) error {
	log.Println("Starting prime")
	return runEngine(parseOptions(options))
}

//todo special thanks to ajhager && shurcooL in the readme for all the code and examples
//todo add alternatives to the readme.md
//todo special thanks to ibon tolosana && jon valdes
//inspired in pixi.js 2 and engi
