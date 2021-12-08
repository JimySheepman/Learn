package main

import "fmt"

func main() {
	type Player struct {
		name      string
		hitpoints int
		stamina   int
	}
	var p1 Player
	var p2 Player2
	p1.name = "Merlins"
	p2.name = "Jimmy"
	p2.hitpoints = 100
	p2.stamina = 80
	fmt.Println("p1 name:", p1.name)
	fmt.Println("p2 name:", p2.name)
	playerLog(p2)

	var p3 Player3
	p3.name = "Jimmy Sheepman"
	playerLog2(&p3)
}

type Player2 struct {
	name      string
	hitpoints int
	stamina   int
}
type Player3 struct {
	name      string
	hitpoints int
	stamina   int
}

func playerLog(p Player2) {
	fmt.Println("Player name:", p.name)
	fmt.Println("Player hitpoints:", p.hitpoints)
	fmt.Println("Player stamina:", p.stamina)
}
func playerLog2(p *Player3) {
	fmt.Println("Player name:", p.name)
	fmt.Println("Player hitpoints:", p.hitpoints)
	fmt.Println("Player stamina:", p.stamina)
}
