// Created by nazarigonzalez on 2/1/17.

package prime

import "log"

type Game struct {
	A, B int
}

func (g *Game) Init() {
	log.Println("Game init")
}
