package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

const (
	ServerAddress = "localhost:8080"
)

func main() {
	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		fmt.Printf("Error de conexión al servidor: %v\n", err)
		return
	}
	defer conn.Close()

	br := bufio.NewReader(conn)
	jStr, _ := br.ReadString('\n')
	jStr = jStr[:len(jStr)-1]
	numPlayer, _ := strconv.Atoi(jStr)
	msg, _ := br.ReadString('\n')
	fmt.Printf(msg)
	fmt.Printf("Soy el jugador %d...\n", numPlayer)
	fmt.Println("Esperando a que el juego comience...\n")

	// Implementa la lógica para esperar a que el juego comience.

	fmt.Println("INICIO DEL JUEGO")
	for {
		// Implementa la lógica para hacer movimientos en el juego.
	}
}
