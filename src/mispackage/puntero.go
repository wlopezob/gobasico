package mispackage

import "fmt"

func MainPuntero() {
	// a es de tipo int
	a := 25
	fmt.Println("Valor de a:", a)
	fmt.Println("Valor del puntero a:", &a)

	b := &a
	// se imprime el contenido de b
	fmt.Println("Valor de b:", *b)

	*b = 32
	fmt.Println("Valor de b:", *b)
	fmt.Println("Valor de b:", a)

	//incremento a
	a++
	fmt.Println("Valor valor que apunta b:", *b)

	if b != nil {
		fmt.Println("b es diferente de nil")
	}

	c := &a
	if b == c {
		fmt.Println("b y c son iguales")
	}

	// otra forma de obtener un puntero
	d := new(int)
	fmt.Println("Direccion de d", d)
	fmt.Println("Valor de d", *d)

	d = b
	fmt.Println("Direccion de d", d)
	fmt.Println("Valor de d", *d)
	fmt.Println("Valor de a", a)
	fmt.Println("Valor de b", *b)
	fmt.Println("Valor de c", *c)

	numero := 1
	incrementar(&numero)
	fmt.Println("Valor de la variable numero: ", numero)
}

func incrementar(numero *int) {
	*numero++
	fmt.Println("Valor numero dentro de la funcion incrementar:", *numero)
}
