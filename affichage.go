package hangman

import (
	"fmt"
	"github.com/fatih/color"
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
