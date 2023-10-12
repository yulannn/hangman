package hangman

import (
	"fmt"
)

func AfficherMot(motSecret string, lettresDevinées []string) {
	for _, lettre := range motSecret {
		lettreStr := string(lettre)
		if ContainsString(lettresDevinées, lettreStr) {
			fmt.Printf("%s ", lettreStr)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func ContainsString(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

func IsLowercaseLetter(lettre string) bool {
	return lettre >= "a" && lettre <= "z"
}
