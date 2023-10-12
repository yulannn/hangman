package main

import (
	"fmt"
	"hangman"
	"strings"
)

func main() {
	lettresDevinées := []string{}
	chance_max := 10

	motSecret, err := hangman.RandomWord()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération du mot aléatoire : %v\n", err)
		return
	}
	hangman.Affiche()
	fmt.Println("Bienvenue au jeu du pendu !")

	for chance := 0; chance < chance_max; chance++ {
		var lettre string
		fmt.Print("Entrez une lettre : ")
		fmt.Scanln(&lettre)

		lettre = strings.ToLower(lettre)
		if len(lettre) != 1 || !isLowercaseLetter(lettre) {
			fmt.Println("Veuillez entrer une seule lettre en minuscules.")
			continue
		}

		if hangman.ContainsString(lettresDevinées, lettre) {
			fmt.Println("Vous avez déjà deviné cette lettre.")
			continue
		}

		lettresDevinées = append(lettresDevinées, lettre)

		hangman.AfficherMot(motSecret, lettresDevinées)

		if chance == chance_max-1 {
			fmt.Println("VOUS AVEZ PERDU")
			fmt.Printf("LE MOT ETAIT %s", motSecret)
		}

	}
}

func isLowercaseLetter(lettre string) bool {
	return lettre >= "a" && lettre <= "z"
}
