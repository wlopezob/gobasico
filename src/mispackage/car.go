package mispackage

import "fmt"

// CardPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// cardPrivate Car con acceso privado
type cardPrivate struct {
	brand string
	year  int
}

// funcion para imprimir mensaje
func PrintMessage(cad string) {
	fmt.Printf("Hola %s\n", cad)
}
