package main

import (
	"fmt"
)

type estudiante struct {
	nombre string
	edad   int
}

type profesor string

type saludador interface {
	saludar(chan string)
}

func (e estudiante) saludar(c chan string) {
	c <- fmt.Sprintf("Hola profesor, soy %s", e.nombre)
}

func (p profesor) saludar(c chan string) {
	c <- fmt.Sprintf("Hola alumno, soy %s", p)
}

func enviarSaludos(s saludador, c chan string) {
	fmt.Printf("Tipo: %T, valor: %v\n", s, s)
	s.saludar(c)
}

// START OMIT
func main() {
	value := "Hola mundo!"
	defer func(s string) { fmt.Println(s) }(value)
	var saludos chan string = make(chan string)
	profesores := []profesor{"Rosita", "Leandro"}
	profesores = append(profesores, "Ariel")
	profesores = profesores[1:]
	for _, profesor := range profesores {
		fmt.Printf("Profesor: %s\n", profesor)
	}
	franco := estudiante{nombre: "Franco"}
	go enviarSaludos(franco, saludos)
	go enviarSaludos(profesores[0], saludos)
	fmt.Println(<-saludos)
	fmt.Println(<-saludos)
	value = "Chau mundo!"
}

// END OMIT
