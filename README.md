# TA3 Modified Ludo
Ludo modificado en lenguaje GO para el curso de programación concurrente y distribuída.

## Descripción del Juego
### Elementos del Laberinto
* El **tablero** está dividido en casillas.
* Cada **casilla** puede tener caminos u obstáculos.
* Los **obstáculos** pueden ser paredes, trampas o criaturas que bloquean el paso de los personajes.
* Cada **jugador** tiene cuatro personajes.
* Cada **personaje** comienza en un punto de partida específico.
### Turnos y Movimientos
En cada turno, los jugadores:
* Lanzan 3 dados:
  * 2 dados normales (del 1 al 6)
  * 1 dado con operación (sumar o restar)
* Mueven solo 1 de sus personajes:
  * Debe avanzar el valor del primer dado +/- (usando el resultado del dado con operación) el valor del segundo dado.
  * Si le toca avanzar hacia una casilla con obstáculo, el jugador pierde el turno y continua el siguiente jugador.
### Objetivo
* Llevar a los cuatro personajes desde los puntos de partida hasta la meta en el menor número de turnos posible.
* El primer jugador en llevar a todos sus personajes a la meta, gana. 

## Concurrencia
### Modificaciones y Uso de Canales
* Los jugadores y el tablero están representados como entidades concurrentes separadas que se comunican a través de canales.
* Cada jugador tiene su propio canal de comunicación con el tablero del juego para enviar movimientos y recibir actualizaciones del estado del juego. 
### Tareas a Implementar
* Implementación de Canales:
  * Cada jugador tiene un canal de comunicación bidireccional con el tablero.
* Lógica del Movimiento:
  * Manejar las interacciones con obstáculos.
