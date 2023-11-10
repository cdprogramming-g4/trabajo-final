package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
)

const (
	ServerAddress = "localhost:8080"
)

type LocalPlayer struct {
	ID         uint
	characters []int
	boardRef   []uint8
	boardSize  uint
	missTurn   bool
	gameOver   bool
	conn       net.Conn
}
type Best struct {
	index    int
	position int
}

func NextMovement(p *LocalPlayer, board []uint8, move int) Best {
	var wg sync.WaitGroup
	bestChann := make(chan Best, len(p.characters)+1)
	bestChann <- Best{-1, -12}

	for i, posChar := range p.characters {
		if posChar == int(p.boardSize) {
			continue
		}
		wg.Add(1)
		go func(index, position int) {
			defer wg.Done()
			newPos := position + move
			fmt.Println("\tPersonaje", index+1, "avanzaría de", position, "a", newPos)
			// Se dirige a un camino libre
			if newPos > 0 && newPos < int(p.boardSize) && board[newPos] == 0 {
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
		best.index = rand.Intn(len(p.characters))
		best.position = p.characters[best.index] + move
	}
	fmt.Println("\tEs mejor mover el personaje", best.index+1)
	return best
}

func (p *LocalPlayer) PlayTurn() {
	// depends on dices
	var move int = 0
	best := NextMovement(p, p.boardRef, move)
	charIndex := best.index
	newPos := best.position

}

func main() {
	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		fmt.Printf("Error de conexión al servidor: %v\n", err)
		return
	}
	defer conn.Close()

	var boardRef []uint8
	br := bufio.NewReader(conn)
	str, _ := br.ReadString('\n')
	parts := strings.Split(str, " ")
	numPlayer, _ := strconv.Atoi(parts[0])
	numCharac, _ := strconv.Atoi(parts[1])
	json.Unmarshal([]byte(parts[2]), &boardRef)
	bSize := len(boardRef)

	localPlayer := &LocalPlayer{
		ID:         uint(numPlayer),
		characters: make([]int, numCharac),
		boardRef:   boardRef,
		boardSize:  uint(bSize),
		missTurn:   false,
		gameOver:   false,
		conn:       conn,
	}

	msg, _ := br.ReadString('\n')
	fmt.Printf(msg)
	fmt.Printf("Soy el jugador %d...\n", localPlayer.ID)
	fmt.Println("Esperando a que el juego comience...\n")

	// Esperar a que el servidor mande la señal para que inicie el juego.
	msg, _ = br.ReadString('\n')
	fmt.Println(msg)

	for {
		// Esperar a que el servidor le de señal de turno
		//...
		localPlayer.PlayTurn()
		// Implementa la lógica para hacer movimientos en el juego.
	}
}
