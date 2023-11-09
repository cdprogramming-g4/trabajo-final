package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	ServerAddress = "localhost:8080"
)

type LocalPlayer struct {
	ID         uint
	characters []int
	boardRef   []uint8
	boardSize  uint
	conn       net.Conn
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
	json.Unmarshal([]byte(strings.TrimSpace(parts[2])), &boardRef)
	bSize := len(boardRef)

	localPlayer := &LocalPlayer{
		ID:         uint(numPlayer),
		characters: make([]int, numCharac),
		boardRef:   boardRef,
		boardSize:  uint(bSize),
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

		// Implementa la lógica para hacer movimientos en el juego.
	}
}
