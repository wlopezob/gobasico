package mispackage

import "fmt"

//Definimos un nuevo tipo
type Dinero int

func (dinero Dinero) String() string {
	return fmt.Sprintf("$%d", dinero)
}

func MainType() {
	var sueldo Dinero
	sueldo = 1000
	fmt.Println("Sueldo:", sueldo)

	aumento := 10000
	sueldo += Dinero(aumento)

	fmt.Println("Sueldo + aumento", sueldo)
}
