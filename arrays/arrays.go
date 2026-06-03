package main;
import "fmt";

func main() {
	fmt.Println("Arreglos");
	fmt.Println("");

	// Declaración
	// var myArray [3]string // 3 cadenas de texto literales

	// Declaración e inicialización
	 var myArray = [3]string{"First","Second","Third"};

	// Cuente los items por nosotros
	// var myArray = [...]string{"First", "Second", "Third"}

	// Estableciendo un nuevo valor a una posición del arreglo existente
	myArray[2] = "Another"
	fmt.Println(myArray[2]) // Accediendo al elemento

	// Viendo la longitud del arreglo
	fmt.Println(len(myArray))

	// Copiando el arreglo original
	anotherArray := myArray
	fmt.Println("Viendo la copia del arreglo: ", anotherArray)

	// Ejemplo:
	// var myArray = [3]string{"First","Second","Third"};
	myArrayCopy := myArray
	myArray[2] = "Another"
	
	fmt.Println("Observando que el elemento sea Another: ",  myArray[2])
	fmt.Println("Observando que el elemento del nuevo arreglo sea Third: " + myArrayCopy[2])

}