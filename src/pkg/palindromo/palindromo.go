package palindromo

import "strconv"

type VerificaPalindromo struct{}

func (vp *VerificaPalindromo) EsPalindromo(numero int) bool {
	numeroStr := strconv.Itoa(numero)
	runeSlice := []rune(numeroStr)

	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}

	numeroRevertido := string(runeSlice)
	return numeroStr == numeroRevertido
}
