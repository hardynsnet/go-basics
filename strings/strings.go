package main
import "fmt"
import (
	"strings"
)

func main() {
	var name = "test"
	fmt.Println(name)

	// probando la funcion len
	fmt.Println(len(name))

	// obteniendo el primer indice de la variable 'name'
	fmt.Println(name[0])

	// obteniendo el segundo indice de la variable 'name'
	fmt.Println(name[1])

	// obteniendo porciones de las cadenas usando una sintaxis
	fmt.Println(name[0:2])
	fmt.Println(name[:2])
	fmt.Println(name[2:])

	// Haciendo una copia de la cadena
	var newstring = name[:]
	fmt.Println(newstring)

	// Asignando la cadena a otra variable 
	var first = "test"
	var second = first // -> test

	first = "another test"
	fmt.Println(first)
	fmt.Println(second)

	// Concatenando
	var word = first + " " + second
	fmt.Println(word)

	var tienePrefijo = strings.HasPrefix("test","te")
	fmt.Println(tienePrefijo)

}