package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var motrandom string

func RandomWord() {
	file, err := os.Open("Dictionnaire.txt")
	if err != nil {
		fmt.Println("Impossible d'ouvrir le fichier de mots.")
		return
	}
	defer file.Close()

	var mots []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur de lecture du fichier de mots.")
		return
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(mots))
	motrandom = mots[randomIndex]

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Print(" Mot à deviner :") 
	Hide()
	fmt.Println("")
	fmt.Println("")
}
