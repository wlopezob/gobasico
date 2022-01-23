package mispackage

import (
	"fmt"
	"time"
)

func enviarPing(c chan string) {
	cont := 0
	for {
		//enviando valores a traves del channel
		cont++
		c <- fmt.Sprintf("Ping %d", cont)
	}
}
func imprimirPing(c chan string) {
	var contador int
	for {
		//Recibiendo valores a traves de un canal
		contador++
		fmt.Println(<-c, " ", contador)
		time.Sleep(time.Second * 1)
	}
}
func Channel01Main() {
	//creamos un channel
	ch := make(chan string)

	//llamamos a la funcion como goroutines
	go enviarPing(ch)
	go imprimirPing(ch)

	//ecaneamos la entrada de dato para que finalice main
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

}
