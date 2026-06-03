# Arrays con GO
Los arreglos son una secuencia de elementos de un solo tipo

Definimos un array de esta manera:

```
var myArray [3]string // un arreglo de 3 cadenas de texto 
```

Se puede inicializar un arreglo con valores así:

```
var myArray = [3]string{"First","Second","Third"}
```

En este caso Go también podemos dejar que Go cuente los elementos por nosotros
```
var myArray = [...]string{"First","Second","Third"}
```

Un arreglo puede solo contener valores del mismo tipo
El arreglo no puede ser redimensionado, se tiene que definir la longitud de un array explicitamente. Es paete del tipo de un arreglo
Tambien, no puedes usar una variable para definir la longitud del arreglo.

Dado esta limitación, los arreglos son raramente usados directamente en Go, y en vez de eso se usan <b>slices</b>. Slices usan los arrays por debajo, y es necesario conocerlos para saber como trabajan.

Puedes acceder a un elemento de un arreglo usando los corchetes_

```
myArray[0]
myArray[1]
```
Se puede definir un nuevo valor para una posición especifica de un arreglo

```
len(myArray)
```

Los arreglos son tipos de un valor. Lo que significa copiar un arreglo.
```
anotherArray := myArray
```
O pasar un arreglo a una funcion, o retornarla de una función, crear una copia del arreglo original.

Esto es diferente de otros lenguajes

Hagamos un ejemplo simple, donde asignamos un nuevo valor a un elemento de un arreglo después de copiarlo.

```
var myArray = [3]string{"First", "Second", "Third"}
myArrayCopy := myArray
myArray[2] = "Another"
myArray[2] // "Another"
myArrayCopy[2] //"Third"
```

Recuerda que solo puedes añadir un solo tipo de elemento en un arreglo, asi que poner ```myArray[2] = 2``` dará un error.
