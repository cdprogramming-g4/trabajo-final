package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

const (
	BoardSize     = 64
	NumPlayers    = 4
	NumCharacters = 4
	NumObstacles  = 10
)

type Player struct {
	ID         uint
	characters []int //Positions
	missTurn   chan bool
}

type BoardSquare uint8

const (
	PATH     BoardSquare = 0
	WALL     BoardSquare = 1
	TRAP     BoardSquare = 2
	CREATURE BoardSquare = 3
)

type Game struct {
	board      []BoardSquare
	gameOver   bool
	turnSignal chan int
}

type Best struct {
	index    int
	position int
}

func NextMovement(p *Player, g *Game, move int) Best {
	var wg sync.WaitGroup
	bestChann := make(chan Best, NumCharacters+1)
	bestChann <- Best{-1, -12}

	for i, posChar := range p.characters {
		if posChar == BoardSize {
			continue
		}
		wg.Add(1)
		go func(index, position int) {
			defer wg.Done()
			newPos := position + move
			fmt.Println("\tPersonaje", index+1, "avanzaría de", position, "a", newPos)
			// Se dirige a un camino libre
			if newPos > 0 && newPos < BoardSize && g.board[newPos] == PATH {
				best := <-bestChann
				// Ha avanzado más
				if newPos > best.position {
					fmt.Println("\tPersonaje", index+1, "en posición", newPos, "es mejor que el personaje", best.index+1)
					current := Best{index, newPos}
					bestChann <- current
				} else {
					bestChann <- best
				}
			}
		}(i, posChar)
	}
	wg.Wait()

	best := <-bestChann
	close(bestChann)
	// Elegir aleatorio si no pudo encontrar uno ideal
	if best.index == -1 {
		best.index = rand.Intn(NumCharacters)
		best.position = p.characters[best.index] + move
	}
	fmt.Println("\tEs mejor mover el personaje", best.index+1)
	return best
}

func (p *Player) Play(g *Game) {
	for true {
		_, ok := <-g.turnSignal
		if !ok {
			return
		}
		fmt.Printf("Turno del jugador %d\n", p.ID+1)

		if <-p.missTurn {
			g.turnSignal <- 1
			p.missTurn <- false
			fmt.Printf("El jugador %d ha perdido su turno\n", p.ID+1)
			continue
		}

		// Lanzar dados
		dice1 := rand.Intn(6) + 1
		dice2 := rand.Intn(6) + 1
		var move int = 0

		if operator := rand.Intn(2); operator == 0 {
			move = dice1 + dice2
			fmt.Printf("Dado: %d + %d\n", dice1, dice2)
		} else {
			move = dice1 - dice2
			fmt.Printf("Dado: %d - %d\n", dice1, dice2)
		}

		best := NextMovement(p, g, move)
		charIndex := best.index
		newPos := best.position

		if newPos < 0 {
			p.characters[charIndex] = 0
			p.missTurn <- false
			fmt.Printf("Personaje %d del jugador %d regresa al inicio\n", charIndex+1, p.ID+1)
		} else if newPos >= BoardSize {
			p.characters[charIndex] = BoardSize
			p.missTurn <- false
			fmt.Printf("El personaje %d del jugador %d llegó a la meta\n", charIndex+1, p.ID+1)
		} else if g.board[newPos] != PATH {
			// p.characters[charIndex] = newPos
			p.missTurn <- true
			fmt.Printf("El personaje %d del jugador %d cayó en un obstáculo, pierde el turno\n", charIndex+1, p.ID+1)
		} else {
			p.characters[charIndex] = newPos
			p.missTurn <- false
			fmt.Printf("El jugador %d avanzó/retrocedió el personaje %d a la casilla %d\n", p.ID+1, charIndex+1, p.characters[charIndex])
		}

		if isWinner(p.characters) {
			fmt.Printf("El jugador %d ganó la partida.\nFIN DEL JUEGO.\n", p.ID+1)
			g.gameOver = true
			close(g.turnSignal)
			return
		}

		if !g.gameOver {
			g.turnSignal <- 1
		}
	}
}

func isWinner(positions []int) bool {
	for _, pos := range positions {
		if pos < BoardSize {
			return false
		}
	}
	return true
}

func PrintBoard(board []BoardSquare, players []*Player) {
	width := int(math.Sqrt(BoardSize))
	var occupied BoardSquare = 9
	for _, p := range players {
		for _, c := range p.characters {
			if c < BoardSize {
				board[c] = occupied
			}
		}
	}
	for i := 0; i < BoardSize; i++ {
		fmt.Printf("%d ", board[i])
		if (i+1)%width == 0 {
			fmt.Printf("\n")
		}
	}
}

func main() {
	game := &Game{
		board:      make([]BoardSquare, BoardSize),
		turnSignal: make(chan int, 1),
	}

	for i := 0; i < NumObstacles; i++ {
		min := len(game.board) / NumObstacles * i
		max := min + NumObstacles
		obsPos := rand.Intn(max-min) + min
		game.board[obsPos] = BoardSquare(rand.Intn(3) + 1)
	}
	game.turnSignal <- 1

	PrintBoard(game.board, []*Player{})

	players := make([]*Player, NumPlayers)
	var wg sync.WaitGroup

	for i := uint(0); i < NumPlayers; i++ {
		players[i] = &Player{
			ID:         i,
			characters: make([]int, NumCharacters),
			missTurn:   make(chan bool, 1),
		}
		// Inicializar la pérdida del turno en falso
		players[i].missTurn <- false

		wg.Add(1)
		go func(p *Player) {
			defer wg.Done()
			p.Play(game)
		}(players[i])
	}

	fmt.Println("INICIO DEL JUEGO")
	wg.Wait()
	PrintBoard(game.board, players)
}
