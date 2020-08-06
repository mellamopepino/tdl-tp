// START OMIT

package main

import (
	"fmt"
	"math"
)

const pi float32 = math.Pi

func main() {
	var saludo string = "Hola mundo!"
	func() {
		fmt.Println(saludo)
	}()
	alumno := estudiante{persona{nombre: "Franco"}, 100615}
	fmt.Printf("Tipo: %T, Valor: %v\n", alumno, alumno)
	alumno.estudiar()
	fmt.Println(alumno.saludar("espectadores"))
}

// END OMIT

type persona struct {
	nombre string
	edad   int
}

type estudiante struct {
	persona
	padron int
}

func (e estudiante) estudiar() {
	fmt.Printf("Soy %s, navegando en stackoverflow\n", e.nombre)
}

type estudiador interface {
	estudiar()
}

func (p persona) saludar(nombre string) string {
	return fmt.Sprintf("Hola %s, soy %s.\n", nombre, p.nombre)
}
