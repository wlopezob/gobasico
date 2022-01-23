package mispackage

import (
	"fmt"
	"time"
)

func generarNumeros(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}

func elevarAlCuadrado(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func imprimir(in <-chan int) {
	for res := range in {
		fmt.Println(res)
		time.Sleep(1 * time.Second)
	}
}
func MainChannel04() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//generamos los numeros
	go generarNumeros(numero)

	//generamos al cuadrado
	go elevarAlCuadrado(numero, cuadrado)

	//imprimimos
	imprimir(cuadrado)
}
