package hangman

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Jeu() {
	fichier, err := os.Create("resultats.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer fichier.Close()
	lettresDevinées := []string{}
	chance_max := 10

	motSecret, err := RandomWord()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération du mot aléatoire : %v\n", err)
		return
	}
	clearConsole()
	color.Blue("Bienvenue au jeu du pendu !")

	for chance := 0; chance < chance_max; {

		switch chance {
		case 1:
			Ascii1()
		case 2:
			Ascii2()
		case 3:
			Ascii3()
		case 4:
			Ascii4()
		case 5:
			Ascii5()
		case 6:
			Ascii6()
		case 7:
			Ascii7()
		case 8:
			Ascii8()
		case 9:
			Ascii9()

		}

		var devinerMotOuLettre string
		color.Cyan("Il vous reste %d tentatives \n", chance_max-chance)
		color.Cyan("[?] Voulez-vous deviner un mot complet ou une lettre ? (m/l): ")
		fmt.Scan(&devinerMotOuLettre)
		clearConsole()
		devinerMotOuLettre = strings.ToLower(devinerMotOuLettre)
		AfficherLettresDevinées(lettresDevinées, motSecret)
		if devinerMotOuLettre == "m" {
			var motDevine string
			fmt.Print("Entrez le mot complet : ")
			fmt.Scan(&motDevine)

			if motDevine == motSecret {
				clearConsole()
				color.Green("Félicitations ! Vous avez deviné le mot : %s\n", motSecret)
				fichier.WriteString("Partie gagnée : " + motSecret + "\n")
				break
			} else {
				clearConsole()
				color.Red("Ce n'est pas le bon mot.")
				chance += 2
				fichier.WriteString("Tentative incorrecte : " + motDevine + "\n")
				if chance >= chance_max {
					clearConsole()
					color.Red("VOUS AVEZ PERDU")
					color.Red("LE MOT ETAIT %s\n", motSecret)
					fichier.WriteString("Partie perdue : " + motSecret + "\n")
					Ascii10()
				}
				continue

			}

		} else if devinerMotOuLettre == "l" {
			var lettre string
			fmt.Print("Entrez une lettre : ")
			fmt.Scan(&lettre)
			lettre = strings.ToLower(lettre)
			fichier.WriteString("Lettre demandée : " + lettre + "\n")

			if len(lettre) != 1 || !IsLowercaseLetter(lettre) {
				fmt.Println("Veuillez entrer une seule lettre en minuscules.")
				continue
			}

			if ContainsString(lettresDevinées, lettre) {
				color.Red("Vous avez déjà deviné cette lettre.")
				fichier.WriteString("Lettre déja choisie \n\n")
				continue
			}

			lettresDevinées = append(lettresDevinées, lettre)

			if !strings.Contains(motSecret, lettre) {
				color.Red("La lettre n'est pas dans le mot.")
				chance += 1
				fichier.WriteString("Lettre incorrecte \n\n")
			} else {
				fichier.WriteString("Lettre correcte \n\n")
			}

			AfficherMot(motSecret, lettresDevinées)
		} else {
			fmt.Println("Répondez avec 'Mot' ou 'Lettre'.")
		}

		if chance >= chance_max {
			clearConsole()
			color.Red("VOUS AVEZ PERDU")
			Ascii10()
			fmt.Printf("LE MOT ETAIT %s\n", motSecret)
			fichier.WriteString("Partie perdue : " + motSecret + "\n")
		}

		if MotDevine(motSecret, lettresDevinées) {
			clearConsole()
			fmt.Println("Félicitations ! Vous avez gagné en devinant toutes les lettres du mot.")
			fichier.WriteString("Partie gagnée, le mot était : " + motSecret + "\n")
			break
		}
	}
}

func MotDevine(motSecret string, lettresDevinées []string) bool {
	for _, lettre := range motSecret {
		if !strings.ContainsRune(strings.Join(lettresDevinées, ""), lettre) {
			return false
		}
	}
	return true
}

func AfficherLettresDevinées(lettresDevinées []string, motSecret string) {
	fmt.Print("Lettres déjà utilisées : ")
	for _, lettre := range lettresDevinées {
		if strings.Contains(motSecret, lettre) {
			color.New(color.FgGreen).Print(lettre)
		} else {
			color.New(color.FgRed).Print(lettre)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}
