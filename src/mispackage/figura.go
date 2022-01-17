package mispackage

import "fmt"

type Figuras2D interface {
	Area() float64
}

func Calcular(f Figuras2D) {
	fmt.Println("Area:", f.Area())
}

type Cuadrado struct {
	Base float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (c Cuadrado) Area() float64 {
	return c.Base * c.Base
}

func (c Rectangulo) Area() float64 {
	return c.Altura * c.Base
}
