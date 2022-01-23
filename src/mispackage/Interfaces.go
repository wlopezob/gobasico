package mispackage

import "fmt"

type Persona struct {
	nombre, email string
	edad          int
}

type Moderador struct {
	Persona
	Foro string
}

type Administrador struct {
	Persona
	Foro string
}

func (m Moderador) AbrirForo() {
	fmt.Println("Abrir foro")
}
func (m Moderador) CerrarForo() {
	fmt.Println("Cerrar foro")
}

func (m Administrador) AbrirForo() {
	fmt.Println("Abrir foro")
}
func (m Administrador) CerrarForo() {
	fmt.Println("Cerrar foro")
}

func (persona Persona) Nombre() string {
	return persona.nombre
}
func (persona Persona) Email() string {
	return persona.email
}

// func PresentarseA(persona Administrador) {
// 	fmt.Println("Nombre:", persona.Nombre())
// 	fmt.Println("Email:", persona.Email())
// }

// func PresentarseM(persona Moderador) {
// 	fmt.Println("Nombre:", persona.Nombre())
// 	fmt.Println("Email:", persona.Email())
// }

// Implementando interfaces

type Usuario interface {
	Nombre() string
	Email() string
}

func Presentarse(usuario Usuario) {
	fmt.Println("Nombre:", usuario.Nombre())
	fmt.Println("Email:", usuario.Email())
}

func MainInterfaces() {
	alejandro := Persona{"alejando", "mail@mail.com", 30}
	Presentarse(alejandro)

	moderador := Moderador{Persona{"moderador", "mail@mail.com", 30}, "Juegos"}
	Presentarse(moderador)
	//PresentarseM(moderador)

	administrador := Administrador{Persona{"administrador", "mail@mail.com", 30}, "PC"}
	Presentarse(administrador)
	//PresentarseA(administrador)

	var i Usuario
	i = moderador

	fmt.Println(i)
	fmt.Println(i.Email())
	fmt.Println(i.Nombre())

	i = administrador

	fmt.Println(i)
	fmt.Println(i.Email())
	fmt.Println(i.Nombre())
}
