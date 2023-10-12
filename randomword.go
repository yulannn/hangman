package hangman

import ("fmt"
"os"
"bufio"
"math/rand"
"time"
)

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
	motrandom := mots[randomIndex]
	fmt.Printf("Mot au hasard : %s\n", motrandom)
}
