# Slices con GO

Un slice es una estructura de datos similar a un arreglo, pero puede cambiar en tamaño.

Internamente los slices usan un arreglo y son una abstracción construida sobre ellos que los hacen más flexibles y útiles.

Podrás usar slices de una manera similar a usar arreglos

Por ejemplo:
Puedes definirlos de manera similar a un arreglo, omitiendo la longitud

```
var mySlice []string // un slice de cadenas de texto
```

Se puede inicializar un slice con valores:
```
var mySlice = []string{"First", "Second", "Third"}
// o
mySlice := []string{"First", "Second", "Third"}
```

Se puede crear un slice vacío de una longitus especifica usando la función ```make()```

```
mySlice := []string{"First","Second","Third"}
newSlice := append(mySlice, "Fourth")
```

Notar que asignamos el resultado de append() a un nuevo slice, de otra manera tendríamos error de complicación, el slice original no es modificado, tenemos uno nuevo

También podemos usar la función ```copy()``` para duplicar un slice, asi que son independientes
```
mySlice := []string{"First","Second","Third"}
newSlice := make([]string, 3)
copy(newSlice, mySlice)
``` 

Si el slice, que copias no tiene suficiente espacio (es decir, es más corto que el original) solo los primeros elementos (hasta que haya espacio) serán copiados.

Se puede inicializar un slice desde un arrar
```
myArray := [3]string{"First","Second","Third"}
mySlice = myArray[:]
```
Multiples slices pueden ser usados en el mismo arreglo como array subyacente 
```
myArray := [3]string{"First","Second","Third"}
mySlice := myArray[:]
mySlice2 := myArray[:]
mySlice[0] = "test"
fmt.Println(mySlice2[0])
```

Estos dos slices, comparten la misma memoria y modificar uno de ellos modifica el arreglo subyacente y causa que el otro slice generado desde el arreglo también sea modificado.

Como los arrays, cada elemento en un slice es almacenado en ubicaciones consecutivas de memoria.

Si necesitas realizar operaciones con los slices, puedes solicitar más de la capacidad inicialmente requerida, asi que si necesitas más espacio, el espacio estará disponible para lectura (en vez de encontrar y mover el slice a una nueva ubicación de memoria con más espacio para crecer y disponerse).

Podemos especificar la capacidad añadiendo un tercer parámetro a la función ```make()```

Ejemplo:
```
newSlice := make([]string, 0, 10)
// un slice vacio con capacidad para 10 elementos
```

Asi como los strings, puedes obtener una porción de un slice usando esta sintaxis:
```
mySlice := []string{"First", "Second", "Third"}
newSlice := mySlice[:2] // obtiene los primeros dos elementos
newSlice2 := mySlice[2:] // ignora los dos primeros elementos
newSlice3 := mySlice[1:3] // nuevo slice con elementos en una posición
```
