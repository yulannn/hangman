package hangman

import (
	"fmt"
	"strings"
)

func Jeu() {
	lettresDevinées := []string{}
	chance_max := 10

	motSecret, err := RandomWord()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération du mot aléatoire : %v\n", err)
		return
	}
	fmt.Println("Bienvenue au jeu du pendu !")

	for chance := 0; chance < chance_max; chance++ {
		var devinerMotOuLettre string
		fmt.Printf("Il vous reste %d tentatives \n", chance_max)
		fmt.Print("Voulez-vous deviner un mot complet ou une lettre ? (Mot/Lettre): ")
		fmt.Scanln(&devinerMotOuLettre)
		devinerMotOuLettre = strings.ToLower(devinerMotOuLettre)

		if devinerMotOuLettre == "mot" {
			var motDevine string
			fmt.Print("Entrez le mot complet : ")
			fmt.Scanln(&motDevine)

			if motDevine == motSecret {
				fmt.Printf("Félicitations ! Vous avez deviné le mot : %s\n", motSecret)
				break
			} else {
				fmt.Println("Ce n'est pas le bon mot.")
				chance_max -= 2
				continue
			}
		} else if devinerMotOuLettre == "lettre" {
			var lettre string
			fmt.Print("Entrez une lettre : ")
			fmt.Scanln(&lettre)

			lettre = strings.ToLower(lettre)
			if len(lettre) != 1 || !IsLowercaseLetter(lettre) {
				fmt.Println("Veuillez entrer une seule lettre en minuscules.")
				continue
			}

			if ContainsString(lettresDevinées, lettre) {
				fmt.Println("Vous avez déjà deviné cette lettre.")
				continue
			}

			lettresDevinées = append(lettresDevinées, lettre)

			if !strings.Contains(motSecret, lettre) {
				fmt.Println("La lettre n'est pas dans le mot.")
				chance_max--
			}

			AfficherMot(motSecret, lettresDevinées)
		} else {
			fmt.Println("Répondez avec 'Mot' ou 'Lettre'.")
		}

		if chance == chance_max-1 {
			fmt.Println("VOUS AVEZ PERDU")
			fmt.Printf("LE MOT ETAIT %s\n", motSecret)
		}

		// Vérifier si toutes les lettres ont été devinées
		if MotDevine(motSecret, lettresDevinées) {
			fmt.Println("Félicitations ! Vous avez gagné en devinant toutes les lettres du mot.")
			break
		}
	}
}

// Vérifier si toutes les lettres du mot ont été devinées
func MotDevine(motSecret string, lettresDevinées []string) bool {
	for _, lettre := range motSecret {
		if !strings.ContainsRune(strings.Join(lettresDevinées, ""), lettre) {
			return false
		}
	}
	return true
}
