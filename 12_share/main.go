package main

import (
	"log"
)

func main() {
	g := NewGame()
	for i := 0; i < 3; i++ {
		g.AddAPlayer()
	}
	for i := 0; i < 4; i++ {
		g.AddBPlayer()
	}

	log.Printf("APlayer:[%d]\n", len(g.APlayers))
	log.Printf("BPlayer:[%d]\n", len(g.BPlayers))
	dressIn := GetDressInstance()
	for k, dress := range dressIn.Dm {
		log.Printf("%s : [%s]\n", k, dress.GetColor())
	}
}
