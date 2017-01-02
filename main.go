// Created by nazarigonzalez on 29/12/16.

package prime

import (
	"log"
)

var engineOptions *PrimeOptions

func Initialize(options *PrimeOptions) error {
	log.Println("Starting prime")

	//extend options with default options
	p, err := runPrime(parseOptions(options))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", p)

	return nil
}

//todo thanks to ajhager && shurcooL in the readme for all the code and examples
