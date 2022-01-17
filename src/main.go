package main

import (
	"fmt"
	"gobasico/src/mispackage"
	"strings"
	"sync"
)

type carro struct {
	brand string
	year  int
}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func sayChannel(text string, c chan<- string) {
	c <- text
}
func main() {
	//a := make(map[int][]string)
	//lsa[0] = []string{"ss", ""}

	//interfaces()
	// punteros()
	//accesibilidad()
	//fMyCar()
	//basic()
	//goWaitGroup()
	//channels()
	channels2()
}

func channels2() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"
	fmt.Println(len(c), cap(c))

	// Range y close
	close(c)
	for message := range c {
		fmt.Println(message)
	}

	//select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}

func message(text string, c chan string) {
	c <- text
}
func channels() {
	c := make(chan string, 1)
	fmt.Println("hello")
	go sayChannel("bye", c)
	fmt.Println(<-c)
}
func goWaitGroup() {
	var wg sync.WaitGroup

	fmt.Println("Hola")

	wg.Add(1)
	go func(text string) {
		fmt.Println(text)
		defer wg.Done()
	}("Adios")

	wg.Add(1)
	go say("Mundo", &wg)

	wg.Wait()
}
func interfaces() {
	myCuadrado := mispackage.Cuadrado{Base: 2}
	myRectangulo := mispackage.Rectangulo{Altura: 2, Base: 3}

	mispackage.Calcular(myCuadrado)
	mispackage.Calcular(myRectangulo)

	// lista de interfaces
	myInterface := []interface{}{"hola", 12, 4.9}
	fmt.Println(myInterface...)

}

func punteros() {
	a := 50
	//acceder a la direccion de la memoria
	b := &a
	fmt.Println(a)
	//acceder al valor de la direccion de la memoria
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := mispackage.Pc{Ram: 16, Disk: 200, Brand: "MSI"}
	fmt.Println(myPC)
	myPC.Ping()
	fmt.Println(myPC)

	myPC.DuplicateRam()
	fmt.Println(myPC)

	myPC.DuplicateRam()
	fmt.Println(myPC)

	fmt.Println(myPC.String())
}
func accesibilidad() {
	var myCar mispackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2022
	mispackage.PrintMessage(myCar.Brand)
}
func fMyCar() {
	myCar := carro{brand: "Ford", year: 202}
	fmt.Println(myCar)

	var otherCar carro
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
func basic() {
	// declaracion de constantes
	const pi float64 = 3.1416
	const pi2 = 3.14
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//declaracion de variables enteras
	base := 12
	fmt.Println(base)
	var altura int = 14
	fmt.Println(altura)
	var area int
	fmt.Println(area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	println(areaCuadrado)

	//operadores matematicos
	x := 10
	y := 50
	//suma
	result := x + y
	fmt.Println("Suma:", result)

	//resta
	result = x - y
	fmt.Println("Resta:", result)

	//multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//division
	result = x / y
	fmt.Println("Division", result)

	//modulo o residuo
	result = x % y
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("X:", x)

	//Decremental
	x--
	fmt.Println("Y:", x)

	//fmt
	helloMessage := "hello"
	worldMessage := "world"
	fmt.Println(helloMessage, worldMessage)

	//Printf, con formato
	name := "william"
	edad := 30
	fmt.Printf("%v tiene %d años\n", name, edad)

	//sprintf
	message := fmt.Sprintf("%v tiene %d años", name, edad)
	fmt.Println(message)

	//tipos de datos
	fmt.Printf("HelloMessage: %T\n", message)
	fmt.Printf("Edad: %T\n", edad)

	//funciones
	normalFunction("1")
	tripeArgument(1, 2, "hola")
	value := returnValue(2)
	fmt.Printf("El valor es %d\n", value)

	value1, value2 := doubleReturn(1)
	println("value1", value1, "value2", value2)

	//variable anonima, sin usar
	value1, _ = doubleReturn(2)
	println("value1", value1, "value2", value2)

	//bucles
	//for conditional
	for i := 0; i <= 10; i++ {
		fmt.Println("numeros:", i)
	}
	fmt.Println("")
	//for while
	counter := 0
	for counter < 10 {
		fmt.Println("counter", counter)
		counter++
	}

	//for ever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 10 {
			break
		}
	}

	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d numero par %d\n", i, par)
	}

	conditional(4, 2)
	multipleConditional(2)
	mDefer()
	arrays()
	mMap()
}
func normalFunction(message string) {
	fmt.Printf("Hola mundo %s\n", message)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func conditional(a, b int) {
	if a == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if a%b == 0 {
		fmt.Printf("%d divisible por %d\n", a, b)
	}

	if a > b || a-b > 0 {
		fmt.Printf("%d es mayor %d\n", a, b)
	}
}

func multipleConditional(a int) {
	var result bool
	switch a % 2 {
	case 0:
		result = true
	default:
		result = false

	}
	if result {
		fmt.Printf("%d es numero par\n", a)
	} else {
		fmt.Printf("no %d es numero par\n", a)
	}
}

func mDefer() {
	defer fmt.Println("hola")
	fmt.Println("mundo")
}

func arrays() {
	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//methos in slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	cadenas := []string{"hola", "como", "estas"}
	for _, valor := range cadenas {
		fmt.Println(valor)
	}
	for i := range cadenas {
		fmt.Println(i)
	}

	isPalindro("Oro")
}

func isPalindro(cad1 string) {
	cad1 = strings.ToUpper(cad1)
	var textReverse string
	for i := len(cad1) - 1; i >= 0; i-- {
		textReverse += string(cad1[i])
	}
	if cad1 == textReverse {
		fmt.Printf("el %s texto es palindro\n", cad1)
	} else {
		fmt.Printf("el %s texto no es palindro\n", cad1)
	}
}

func mMap() {
	m := make(map[string]int)
	m["jose"] = 14
	m["pepito"] = 20
	fmt.Println(m)

	//recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	valor, ok := m["jose"]
	fmt.Println(valor, ok)
}
