package main

import ("hangman"
"fmt")

func main() {
	for {
		hangman.Affiche()
		if !hangman.Rejouer() {
			break
		}
	}
	fmt.Println("Merci d'avoir joué ! Au revoir.")
}
