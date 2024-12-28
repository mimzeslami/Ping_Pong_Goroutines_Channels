
# Ping Pong Game with Goroutines and Channels

This project demonstrates a simple **Ping Pong game** implemented in Go, using **goroutines** and **channels** to simulate concurrency. The program showcases how goroutines communicate and synchronize using channels in a real-world scenario.

---

## Features
- Simulates a Ping Pong game with dynamic ball movement.
- Demonstrates Go's concurrency model using **goroutines** and **channels**.
- Includes a command-line interface (CLI) for user interaction.
- Clears the screen dynamically to visualize the ball moving between players.

---

## Concepts Used
### Goroutines
- Goroutines are lightweight threads managed by the Go runtime.
- In this game, two goroutines (`ping` and `pong`) run concurrently, simulating the players in the Ping Pong game.

### Channels
- Channels provide a way for goroutines to communicate with each other.
- A **Ping** signal is sent from one goroutine to another via a channel, and the other goroutine responds with a **Pong** signal.

---

## How It Works
1. **Start the Game**: 
   - The user specifies the number of Ping Pong exchanges.
2. **Signal Flow**:
   - A `pingCh` channel signals the `Ping` goroutine.
   - A `pongCh` channel signals the `Pong` goroutine.
3. **Dynamic Ball Movement**:
   - The ball (`🏓`) moves visually between two players in the terminal.

---

## How to Run
1. Clone the repository or copy the source code into a file named `main.go`.
2. Build and run the program using:
   ```bash
   go run main.go
   ```
3. Follow the prompts in the terminal:
   - Choose to start the game or exit.
   - Enter the number of Ping Pong exchanges.

---

## Example Output

### Before Starting
```
=== Ping Pong Game ===
1. Start Game
2. Exit
Enter your choice: 1
Enter the number of Ping-Pong exchanges: 5
Starting the game...
```

### During the Game
```
=== Ping Pong Game ===

Player 1 [Ping]  🏓                                Player 2 [Pong]
```
Then alternates to:
```
=== Ping Pong Game ===

Player 1 [Ping]                                🏓  Player 2 [Pong]
```

### Game Over
```
Ping
Pong
Ping
Pong
Ping
Pong
Ping
Pong
Ping
Pong
Game over!
```

---

## Learning Objectives
This project is ideal for learning:
- How to create and manage **goroutines** in Go.
- How to use **channels** to synchronize communication between goroutines.
- Implementing concurrency in a simple and fun project.

---

## Author
Mohammadreza Eslami  
Ping Pong Game with Go’s Concurrency Features
