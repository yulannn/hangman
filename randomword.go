package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func RandomWord() (string, error) {
	file, err := os.Open("Dictionnaire.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var mots []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())
	motRandom := mots[rand.Intn(len(mots))]
	return motRandom, nil
}
