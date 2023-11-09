package main

import (
	"fmt"
	"hangman"
)

func main() {
	for {
		hangman.Affiche()
		if !hangman.Rejouer() {
			break
		}
	}

	fmt.Println("Merci d'avoir jouée ! Au revoir.")
}
