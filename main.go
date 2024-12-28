package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type PingPong interface {
	increment()
	ping(wg *sync.WaitGroup, pingCh, pongCh chan struct{})
	pong(wg *sync.WaitGroup, pingCh, pongCh chan struct{})
}

type Game struct {
	count int
	limit int
	mu    sync.Mutex
}

func (g *Game) increment() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.count++
}

func (g *Game) getCount() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.count
}

func (g *Game) ping(wg *sync.WaitGroup, pingCh, pongCh chan struct{}) {
	defer wg.Done()
	for {
		if g.getCount() >= g.limit {
			return
		}
		<-pingCh // Wait for Ping signal
		g.showBallMovement("Ping")
		fmt.Println("Ping")
		time.Sleep(500 * time.Millisecond) // simulate game
		g.increment()
		pongCh <- struct{}{} // Signal Pong
	}
}

func (g *Game) pong(wg *sync.WaitGroup, pingCh, pongCh chan struct{}) {
	defer wg.Done()
	for {
		if g.getCount() >= g.limit {
			return
		}
		<-pongCh // Wait for Pong signal
		g.showBallMovement("Pong")
		fmt.Println("Pong")
		time.Sleep(500 * time.Millisecond) // simulate game
		g.increment()
		pingCh <- struct{}{} // Signal Ping
	}
}

func (g *Game) showBallMovement(position string) {
	// Clear the console
	fmt.Print("\033[H\033[2J")
	fmt.Println("=== Ping Pong Game ===")
	fmt.Println()

	// Simulate ball movement
	if position == "Ping" {
		fmt.Println("Player 1 [Ping]  🏓                                Player 2 [Pong]")
	} else {
		fmt.Println("Player 1 [Ping]                                🏓  Player 2 [Pong]")
	}

	fmt.Println()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Ping Pong Game ===")
	fmt.Println("1. Start Game")
	fmt.Println("2. Exit")
	fmt.Print("Enter your choice: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "2" {
		fmt.Println("Exiting the game. Goodbye!")
		return
	}

	if choice != "1" {
		fmt.Println("Invalid choice. Exiting.")
		return
	}

	fmt.Print("Enter the number of Ping-Pong exchanges: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	limit, err := strconv.Atoi(input)
	if err != nil || limit <= 0 {
		fmt.Println("Invalid number. Exiting.")
		return
	}

	game := &Game{limit: limit}

	pingCh := make(chan struct{}, 1)
	pongCh := make(chan struct{}, 1)
	var wg sync.WaitGroup

	wg.Add(2) // Two goroutines: Ping and Pong

	// Start Ping and Pong goroutines
	go game.ping(&wg, pingCh, pongCh)
	go game.pong(&wg, pingCh, pongCh)

	fmt.Println("Starting the game...")
	pingCh <- struct{}{} // Signal Ping to start the sequence

	wg.Wait()

	fmt.Println("Game over!")
}
