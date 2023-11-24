package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	NumCharacters = 4
	ServerAddress = "localhost:8080"
)

var NumObstacles int = 1
var NumPlayers int = 2
var BoardSize int = 9
var Board []BoardSquare

type Player struct {
	ID         uint
	characters []int
	missTurn   chan bool
	conn       net.Conn
	buff       *bufio.Reader
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
	winner     uint
	turnSignal chan int
	players    []*Player
}

func (p *Player) Play(g *Game) {
	for true {
		_, ok := <-g.turnSignal
		if !ok {
			fmt.Fprintln(p.conn, "Ganó el jugador "+strconv.Itoa(int(g.winner)))
			return
		}
		fmt.Printf("Turno del jugador %d\n", p.ID)

		if <-p.missTurn {
			g.turnSignal <- 1
			p.missTurn <- false
			fmt.Printf("El jugador %d ha perdido su turno\n", p.ID)
			fmt.Fprintln(p.conn, "miss")
			continue
		}

		// Enviar señal de jugada
		fmt.Fprintln(p.conn, "play")

		// Recibir jugada del jugadador
		str, _ := p.buff.ReadString('\n')
		datos := strings.Split(str, " ")
		switch datos[0] {
		case "move":
			idx, _ := strconv.Atoi(datos[1])
			pos, _ := strconv.Atoi(datos[2])
			p.characters[idx] = pos
			p.missTurn <- false
		case "miss":
			p.missTurn <- true
		}

		time.Sleep(500 * time.Millisecond)

		if isWinner(p.characters) {
			fmt.Printf("El jugador %d ganó la partida.\nFIN DEL JUEGO.\n", p.ID)
			g.winner = p.ID
			g.gameOver = true
			close(g.turnSignal)
			fmt.Fprintln(p.conn, "win")
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
	width := int(math.Sqrt(float64(BoardSize)))
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

type Tag string

const (
	MSG_CONFIG Tag = "config"
)

type MessageRes struct {
	Type Tag `json:"type"`
}

type Config struct {
	NumPlayers   int `json:"numPlayers"`
	NumObstacles int `json:"numObstacles"`
	Size         int `json:"size"`
}
type ArrayData struct {
	Type  string      `json:"type"`
	Array interface{} `json:"array"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir todas las solicitudes de origen
	},
}

func ReadConfig(data []byte) {
	// Decodifica el cuerpo JSON en la estructura Config
	var currConfig Config
	if err := json.Unmarshal(data, &currConfig); err != nil {
		fmt.Println("Error al decodificar JSON", err)
		return
	}

	NumObstacles = currConfig.NumObstacles
	NumObstacles = currConfig.NumObstacles
	BoardSize = currConfig.Size

	fmt.Println("Configuración realizada", NumPlayers, NumObstacles, BoardSize)
}
func CreateBoard() {
	Board = make([]BoardSquare, BoardSize)
	for i := 0; i < NumObstacles; i++ {
		min := BoardSize / NumObstacles * i
		max := min + NumObstacles
		obsPos := rand.Intn(max-min) + min
		Board[obsPos] = BoardSquare(rand.Intn(3) + 1)
	}
}

func SendConnMessage(message []byte, conn *websocket.Conn, messageType int) {
	// Enviar una respuesta al cliente (opcional)
	err := conn.WriteMessage(messageType, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func handler(conn *websocket.Conn) {
	defer conn.Close()

	// cfor:
	for {
		// Leer mensaje desde el cliente
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Imprimir el mensaje recibido
		fmt.Printf("\nMensaje recibido: %s\n", data)

		// Parsear el mensaje
		var common MessageRes
		if err := json.Unmarshal(data, &common); err != nil {
			fmt.Println(err)
			return
		}

		switch common.Type {
		case MSG_CONFIG:
			ReadConfig(data)
			CreateBoard()
			var sliceInt []int
			for _, v := range Board {
				sliceInt = append(sliceInt, int(v))
			}
			msg := ArrayData{
				Type:  "Board",
				Array: sliceInt,
			}
			gBoard, err := json.Marshal(msg)
			if err != nil {
				fmt.Println(err)
				return
			}
			SendConnMessage(gBoard, conn, messageType)
			StartGame()
		default:
			fmt.Println("Tag no identificado")
		}

		responseMessage := []byte("Mensaje recibido con éxito")
		SendConnMessage(responseMessage, conn, messageType)
	}

	fmt.Println("Handle end")
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		go handler(conn)
	})
	fmt.Printf("Servidor escuchando en %s\n", ServerAddress)
	http.ListenAndServe(":8080", nil)
}

func StartGame() {
	game := &Game{
		board:      Board,
		turnSignal: make(chan int, 1),
	}
	game.turnSignal <- 1

	// Codificar el tablero para pasarlo como referencia a cada jugador
	boardRefBytes, _ := json.Marshal(game.board)
	boardRefStr := string(boardRefBytes)

	PrintBoard(game.board, game.players)
	game.players = make([]*Player, NumPlayers)

	listener, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor escuchando en", listener.Addr())
	fmt.Println("Esperando a que se conecten todos los jugadores...")

	for i := 0; i < NumPlayers; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error aceptando la conexión al servidor: %v\n", err)
			return
		}
		// ReadMessage(conn)

		buff := bufio.NewReader(conn)

		game.players[i] = &Player{
			ID:         uint(i + 1),
			characters: make([]int, NumCharacters),
			missTurn:   make(chan bool, 1),
			conn:       conn,
			buff:       buff,
		}
		// Inicializar la pérdida del turno en falso
		game.players[i].missTurn <- false

		fmt.Fprintln(conn, strconv.Itoa(i+1)+" "+strconv.Itoa(NumCharacters)+" "+boardRefStr)
		fmt.Fprintln(conn, "Conectado: Eres el jugador "+strconv.Itoa(i+1))
	}

	fmt.Println("Todos los jugadores se han conectado. Esperando a que comience al juego...")
	fmt.Println("INICIO DEL JUEGO")

	var wg sync.WaitGroup

	for _, player := range game.players {
		wg.Add(1)
		go func(p *Player) {
			defer wg.Done()
			p.Play(game)
		}(player)
		fmt.Fprintln(player.conn, "INICIO DEL JUEGO")
	}

	wg.Wait()
	PrintBoard(game.board, game.players)
}
