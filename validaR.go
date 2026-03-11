package main

import (
	"fmt"
	"unicode"
)

func leer() (bool, bool, bool, bool) {

	tieneMayus := false
	tieneMinus := false
	tieneNumero := false
	tieneSigno := false

	var contraseña string
	fmt.Print("Ingresa la contraseña: ")
	fmt.Scan(&contraseña)

	if len(contraseña) >= 4 && len(contraseña) <= 7 {
		fmt.Println("La contraseña es muy corta")
	}
	if len(contraseña) >= 7 && len(contraseña) <= 10 {
		fmt.Println("La contraseña tiene longitud buena")
	}
	if len(contraseña) > 10 && len(contraseña) <= 15 {
		fmt.Println("El largo de la contraseña es excelente")
	}
	if len(contraseña) > 15 {
		fmt.Println("La contraseña es muy larga")
	}
	for _, c := range contraseña {

		if unicode.IsUpper(c) {
			tieneMayus = true
		}

		if unicode.IsLower(c) {
			tieneMinus = true
		}

		if unicode.IsDigit(c) {
			tieneNumero = true
		}

		if unicode.IsPunct(c) {
			tieneSigno = true
		}

	}
	return tieneMayus, tieneMinus, tieneNumero, tieneSigno
}

func lectura() {
	tieneMayus, tieneMinus, tieneNumero, tieneSigno := leer()

	puntos := 0

	if tieneMayus {
		puntos++
	}
	if tieneMinus {
		puntos++
	}
	if tieneNumero {
		puntos++
	}
	if tieneSigno {
		puntos++
	}

	fmt.Println("Puntaje de seguridad:", puntos)

	// Nivel de seguridad
	switch puntos {

	case 1:
		fmt.Println("Contraseña muy débil")

	case 2:
		fmt.Println("Contraseña débil")

	case 3:
		fmt.Println("Contraseña buena")

	case 4:
		fmt.Println("Contraseña fuerte")

	default:
		fmt.Println("Contraseña muy insegura")
	}

}

func main() {
	lectura()

}
