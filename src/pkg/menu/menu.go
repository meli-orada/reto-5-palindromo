package menu

import (
	"bufio"
	"fmt"
	"strconv"

	"palindromo/src/pkg/palindromo"
)

type Menu struct {
	Scanner *bufio.Scanner
	Vp      *palindromo.VerificaPalindromo
}

func (m *Menu) Show() {
	for {
		// Mostrar el menú
		fmt.Println("=== Menú ===")
		fmt.Println("1. Calcular palíndromo de un número entero.")
		fmt.Println("2. Salir.")
		fmt.Println("Cualquier otra opción retornará al menú.")

		// Leer la opción ingresada por el usuario
		fmt.Print("Ingrese una opción: ")
		m.Scanner.Scan()
		opcionNum := m.Scanner.Text()

		// Convertir la opción a un número entero
		opcionTxt, err := strconv.Atoi(opcionNum)

		if err != nil {
			fmt.Println("Opción inválida. Por favor, ingrese un número entero.")
			continue
		}

		// Evaluar la opción ingresada
		switch opcionTxt {
		case 1:
			fmt.Println("Ingrese un número entero: ")
			m.Scanner.Scan()
			numeroTxt := m.Scanner.Text()

			// Convertir la opción a un número entero
			numeroInt, err := strconv.Atoi(numeroTxt)

			if err != nil {
				fmt.Println("No es un valor entero. Por favor, ingrese un número entero.")
				continue
			}

			if m.Vp.EsPalindromo(numeroInt) {
				fmt.Println("El número es un palíndromo")
			} else {
				fmt.Println("El número no es un palíndromo")
			}
		case 2:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción inválida. Por favor, seleccione una opción válida.")
		}
		fmt.Println()
	}
}
