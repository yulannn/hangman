package main

import "fmt"

func main() {
	var filename string
		fmt.Print("Entrez le nom du fichier: ")
		fmt.Scanln(&filename)
		file, aled := os.Open(filename)
		if aled != nil {
			fmt.Println("Erreur:", aled)
		} else {
			defer file.Close()
			data := make([]byte, 1024)
			for {
				n, aled := file.Read(data)
				if aled == io.EOF {
					break
				}
				if aled != nil {
					fmt.Println("Erreur lecture du fichier", aled)
					break
				}
				fmt.Print(string(data[:n]))
			}
		}
}