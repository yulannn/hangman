package main

<<<<<<< HEAD
import (
	"fmt"
	"hangman"
)
=======
import ("hangman"
"fmt")
>>>>>>> 2c7c71505990ec408297b483ff29a491c392797e

func main() {
	for {
		hangman.Affiche()
		if !hangman.Rejouer() {
			break
		}
	}
<<<<<<< HEAD
	fmt.Println("Merci d'avoir jouÃ© ! Au revoir.")
=======
	fmt.Println("Merci d'avoir joué ! Au revoir.")
>>>>>>> 2c7c71505990ec408297b483ff29a491c392797e
}
