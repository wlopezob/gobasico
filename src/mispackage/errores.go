package mispackage

import (
	"errors"
	"fmt"
)

var (
	//Error usuario no valido ...
	ErrorUsuarioNoValido  = errors.New("El usuario no es valido")
	ErrorUsuarioEnProceso = errors.New("Usuario en proceso de registro")
	ErrorPorDefecto       = errors.New("Algo paso xd...")
)

func baneado(usuario string) (err error) {
	ban := false
	switch usuario {
	case "miguel":
		ban = true
	case "carlos":
		ban = false
	case "alberto":
		ban = false
	case "mario":
		return ErrorUsuarioNoValido
	case "juan":
		return ErrorUsuarioEnProceso
	default:
		return ErrorPorDefecto
	}

	if !ban {
		fmt.Printf("El usuario %s no esta baneado\n", usuario)
	} else {
		fmt.Printf("El usuario %s esta baneado\n", usuario)
	}
	return nil
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error\n", err)
	}
}
func MainError() {
	err := baneado("miguel")
	if err != nil {
		fmt.Println("Error\n", err)
	}

	err = baneado("carlos")
	CheckError(err)
	err = baneado("wil")
	CheckError(err)
}
