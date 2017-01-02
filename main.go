// Created by nazarigonzalez on 29/12/16.

package prime

import (
	"log"
)

func Init(options *PrimeOptions) error {
	log.Println("Starting prime")
	return runEngine(parseOptions(options))
}

//todo thanks to ajhager && shurcooL in the readme for all the code and examples
