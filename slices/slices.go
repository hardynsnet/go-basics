package main;
import "fmt";

func main() {

	// Declaración
	// var mySlice []string;

	// Inicialización con valores
	var mySlice = []string{"First", "Second", "Third"};
	// O también
	mySlice := []string{"First","Second","Third"};

	// Creando un slice vacío con especificación de longitud
	mySlice := make([]string, 3); // un slice con tres cadenas vacías

	// Creando un slice a partir de uno existente
	mySlice := []string{"First", "Second", "Third"};
	newSlice := append(mySlice, "Fourth", "Fifth");

	// Duplicando el slice
	mySlice := []string{"First", "Second", "Third"};
	newSlice := make([]string, 3);
	copy(newSlice, mySlice);

	// Inicializando slice desde un arreglo
	myArray := [3]string{"First", "Second", "Third"}
	mySlice = myArray[:]

	// Array subyacente
	myArray := [3]string{"First", "Second", "Third"}
	mySlice := myArray[:]
	mySlice2 := myArray[:]
	mySlice[0] = "test"
	fmt.Println(mySlice2)

	// Añadiendo capacidad del slice
	newSlice := make([]string, 0, 10) // un slice vacío con capacidad de 10

	// Obteniendo porciones
	mySlice := []string{"First", "Second", "Third"}
	newSlice := mySlice[:2] // obtiene los dos primeros items
	newSlice2 := mySlice[2:] // ignora los dos primeros items
	newSlice3 := mySlice[1:3] // nuevo slice con elementos en posicion 1 a 2
}