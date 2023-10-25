package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	BoardSize     = 64
	NumPlayers    = 4
	NumCharacters = 4
)

type Player struct {
	ID         int
	characters []int //Positions
}

type Game struct {
	board     []int
	obstacles map[int]bool
}

func NextMovement(p *Player, g *Game, move int) []int {
	best := make([]int, 2)
	best[0] = -1
	best[1] = 0
	var wg sync.WaitGroup

	for i, posChar := range p.characters {
		wg.Add(1)
		go func(index, position, move int) {
			defer wg.Done()
			newPos := position + move
			// Se dirige a un camino libre
			if newPos > 0 && newPos < BoardSize &&
				!g.obstacles[newPos] {
				// Ha avanzado más
				if newPos > best[1] {
					best[0] = index
					best[1] = newPos
				}
			}
		}(i, posChar, move)
	}
	wg.Wait()

	// Elegir aleatorio si no pudo encontrar uno ideal
	if best[0] == -1 {
		best[0] = rand.Intn(len(p.characters)) //charIndex
		best[1] = p.characters[best[0]] + move //position
	}
	return best
}

func (p *Player) Play(g *Game) {
	for true {
		fmt.Printf("Turno del jugador %d\n", p.ID+1)

		// Lanzar dados
		dice1 := rand.Intn(6) + 1
		dice2 := rand.Intn(6) + 1
		var move int = 0

		if operator := rand.Intn(2); operator == 0 {
			move = dice1 + dice2
		} else {
			move = dice1 - dice2
		}

		best := NextMovement(p, g, move)
		charIndex := best[0]
		newPos := best[1]

		if newPos > 0 && newPos < BoardSize && !g.obstacles[newPos] {
			p.characters[charIndex] = newPos
			fmt.Printf("El jugador %d avanzó/retrocedió el personaje %d a la casilla %d\n", p.ID+1, charIndex+1, p.characters[charIndex])
		} else {
			fmt.Printf("El personaje %d del jugador %d no puede avanzar, pierde el turno\n", charIndex+1, p.ID+1)
		}

		break
	}
}

func main() {
	game := &Game{
		board: make([]int, BoardSize),
	}

	players := make([]*Player, NumPlayers)

	for i := 0; i < NumPlayers; i++ {
		players[i] = &Player{
			ID:         i,
			characters: make([]int, NumCharacters),
		}
		players[i].Play(game)
	}

}
