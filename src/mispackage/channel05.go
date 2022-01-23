package mispackage

import (
	"fmt"
	"os"
	"time"
)

func leerDatos(out chan<- []byte) {
	for {
		datos := make([]byte, 1024)
		n, _ := os.Stdin.Read(datos)
		if n > 0 {
			out <- datos
		}
	}
}

func MainChannel06() {
	done := time.After(20 * time.Second)
	eco := make(chan []byte)
	go leerDatos(eco)
	for {
		select {
		case datos := <-eco:
			os.Stdout.Write(datos)
		case <-done:
			fmt.Println("Se termino el tiempo")
			os.Exit(0)
		}
	}
}
