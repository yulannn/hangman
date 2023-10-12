package hangman

import "fmt"

func Hide() {
	for i := 0; i < len(motrandom); i++ {
		fmt.Printf("_")
	}
}
