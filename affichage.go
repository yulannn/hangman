package hangman

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Affiche() {
	fmt.Println(" ██░ ██  ▄▄▄       ███▄    █   ▄████  ███▄ ▄███▓ ▄▄▄       ███▄    █ ")
	fmt.Println("▓██░ ██▒▒████▄     ██ ▀█   █  ██▒ ▀█▒▓██▒▀█▀ ██▒▒████▄     ██ ▀█   █ ")
	fmt.Println("▒██▀▀██░▒██  ▀█▄  ▓██  ▀█ ██▒▒██░▄▄▄░▓██    ▓██░▒██  ▀█▄  ▓██  ▀█ ██▒")
	fmt.Println("░▓█ ░██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█  ██▓▒██    ▒██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒")
	fmt.Println("░▓█▒░██▓ ▓█   ▓██▒▒██░   ▓██░░▒▓███▀▒▒██▒   ░██▒ ▓█   ▓██▒▒██░   ▓██░")
	fmt.Println(" ▒ ░░▒░▒ ▒▒   ▓▒█░░ ▒░   ▒ ▒  ░▒   ▒ ░ ▒░   ░  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ ")
	color.Red(" ▒ ░▒░ ░  ▒   ▒▒ ░░ ░░   ░ ▒░  ░   ░ ░  ░      ░  ▒   ▒▒ ░░ ░░   ░ ▒░")
	color.Red(" ░  ░░ ░  ░   ▒      ░   ░ ░ ░ ░   ░ ░      ░     ░   ▒      ░   ░ ░ ")
	color.Red(" ░  ░  ░      ░  ░         ░       ░        ░         ░  ░         ░ ")

	fmt.Println("1. Jouer avec un mot Aléatoire")
	fmt.Println("2. Quitter le jeu")
	
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		RandomWord()
	case 2:
		os.Exit(0)
		default: fmt.Println("Faites un choix valide")
	}
}

func Ascii1() {
	fmt.Println("________________")
}

func Ascii2() {
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}

func Ascii3() {
	fmt.Println(" ____________")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}

func Ascii4() {
	fmt.Println(" ____________")
	fmt.Println("  |          |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}

func Ascii5() {
	fmt.Println(" ____________")
	fmt.Println("  |          |")
	fmt.Println("  |          O")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}

func Ascii6() {
	fmt.Println(" ____________")
	fmt.Println("  |          |")
	fmt.Println("  |          O")
	fmt.Println("  |          |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}

func Ascii7() {
	fmt.Println(" ____________")
	fmt.Println("  |          |")
	fmt.Println("  |          O")
	fmt.Println("  |         /|")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}

func Ascii8() {
	fmt.Println(" ____________")
	fmt.Println("  |          |")
	fmt.Println("  |          O")
	fmt.Println("  |         /|\\")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}

func Ascii9() {
	fmt.Println(" ____________")
	fmt.Println("  |          |")
	fmt.Println("  |          O")
	fmt.Println("  |         /|\\")
	fmt.Println("  |          | ")
	fmt.Println("  |         /")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}

func Ascii10() {
	fmt.Println(" ____________")
	fmt.Println("  |          |")
	fmt.Println("  |          O")
	fmt.Println("  |         /|\\")
	fmt.Println("  |          | ")
	fmt.Println("  |         / \\")
	fmt.Println("  |")
	fmt.Println("  |")
	fmt.Println("__|_____________")
}
