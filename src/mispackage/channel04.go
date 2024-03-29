package mispackage

import (
	"fmt"
	"strconv"
	"time"
)

func enviarMensaje(out chan<- string, numero int) {
	for {
		out <- "Mensaje: " + strconv.Itoa(numero)
		fmt.Println("Enviando mensaje func:", numero)
	}
}

func imprimir01(in <-chan string) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
func MainChannel05() {
	ch := make(chan string, 5)
	for i := 0; i < 5; i++ {
		go enviarMensaje(ch, i)
	}

	imprimir01(ch)
}
