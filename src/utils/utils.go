package utils

import "fmt"

func ClearTerminal() {
	// Limpa a tela e posiciona o cursor no início
	fmt.Print("\033[H\033[2J")

}
