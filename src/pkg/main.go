package main

import (
	"bufio"
	"os"

	"palindromo/src/pkg/menu"
	"palindromo/src/pkg/palindromo"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	vp := &palindromo.VerificaPalindromo{}
	menu := &menu.Menu{
		Scanner: scanner,
		Vp:      vp,
	}
	menu.Show()
}
