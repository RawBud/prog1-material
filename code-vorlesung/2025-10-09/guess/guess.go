package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//my_number := 42
	my_number := rand.Intn(100) + 1 // Zufallszahl zwischen 1 und 100

	for i := 0; i < 3; i++ { // Schleife, die 3 Mal durchläuft
		guess := ReadNumber()
		//if guess == my_number
		if NumberGood(guess, my_number) {
			fmt.Println("Das war richtig!")
			return // Beendet das Programm
		}
		fmt.Printf("%d war nicht korrekt!\n", guess) // Ausgabe mit Formatierung

		fmt.Println("Game Over!")
	}
}

// ReadNumber liefert uns ein int.
func ReadNumber() int {
	var n int //alternativ: n := 0

	fmt.Print("Bitte gib eine Zahl ein: ")
	fmt.Scan(&n)

	return n
}

// NumberGood prüft, ob x gleich einer zufällig
// gewählten Zahl zwischen 1 und 100 ist.
// Liefert true falls x gleich dieser Zufallszahl ist, ansonsten false.
func NumberGood(x int, n int) bool {

	/*if x == my_number {
		return true
	}
	return false*/

	return x == n
}
