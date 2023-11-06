package hangman

import ("fmt"
"strings")

func Rejouer() bool {
	var rejouer string
	fmt.Print("Voulez-vous rejouer ? (Oui/Non): ")
	fmt.Scanln(&rejouer)
	rejouer = strings.ToLower(rejouer)
	return rejouer == "oui"
}
