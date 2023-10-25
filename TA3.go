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
	ID         uint
	characters []int //Positions
}

type BoardSquare uint8

const (
	PATH     BoardSquare = 0
	WALL     BoardSquare = 1
	TRAP     BoardSquare = 2
	CREATURE BoardSquare = 3
)

type Game struct {
	board []BoardSquare
}

type Best struct {
	index    int
	position int
}

func NextMovement(p *Player, g *Game, move int) Best {
	var wg sync.WaitGroup
	var mu sync.Mutex
	n := len(p.characters)
	best := Best{-1, 0}

	for i, posChar := range p.characters {
		wg.Add(1)
		go func(index, position int) {
			defer wg.Done()
			newPos := position + move
			// Se dirige a un camino libre
			if newPos > 0 && newPos < BoardSize &&
				g.board[newPos] == PATH {
				// Ha avanzado más
				mu.Lock()
				if newPos > best.position {
					best = Best{index, newPos}
				}
				mu.Unlock()
			}
		}(i, posChar)
	}
	wg.Wait()

	// Elegir aleatorio si no pudo encontrar uno ideal
	if best.index == -1 {
		best.index = rand.Intn(n)                       //charIndex
		best.position = p.characters[best.index] + move //position
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
		charIndex := best.index
		newPos := best.position

		if newPos > 0 && newPos < BoardSize && g.board[newPos] == PATH {
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
		board: make([]BoardSquare, BoardSize),
	}

	players := make([]*Player, NumPlayers)

	for i := uint(0); i < NumPlayers; i++ {
		players[i] = &Player{
			ID:         i,
			characters: make([]int, NumCharacters),
		}
		players[i].Play(game)
	}

}
