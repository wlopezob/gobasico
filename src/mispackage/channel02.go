package mispackage

import (
	"fmt"
	"time"
)

func MainChannel03() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//generamos los numeros
	go func() {
		for x := 0; x < 5; x++ {
			numero <- x
		}
		close(numero)
	}()

	// elevamos al cuadrado
	/*go func() {
		for {
			x, ok := <-numero
			if !ok { //channel cerrado
				break
			}
			cuadrado <- x * x
		}
		close(cuadrado)
	}()*/
	// elevamos al cuadrado
	go func() {
		for x := range numero {
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	// imprimimos en main
	/*for {
		res, ok := <-cuadrado
		if !ok {
			break
		}
		fmt.Println(res)
		time.Sleep(1 * time.Second)
	}*/
	for res := range cuadrado {
		fmt.Println(res)
		time.Sleep(1 * time.Second)
	}
}
func MainChannel02() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//generamos los numeros
	go func() {
		for x := 0; ; x++ {
			numero <- x
		}
	}()

	//lo elevamos al cuadrado
	go func() {
		for {
			x := <-numero
			cuadrado <- x * x
		}
	}()

	// imprimimos en main
	for {
		fmt.Println(<-cuadrado)
		time.Sleep(1 * time.Second)
	}

}
