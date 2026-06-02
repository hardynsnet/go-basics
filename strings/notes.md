# Strings en GO
Una cadena en Go is una seecuencia de valores de bytes.
Podemos definir una cadena usando esta sintaxis:

<b>var name = "test"</b>

Es importante notar que en otros lenguajes, las cadenas se definen solo usando comillas dobles, no comillas simples.

Para obtener la longitud de una cadena usamos la función integrada <b>'len()'</b>, ejemplo:
* len(name) // Devuelve que tiene 4 letras

Y también, puedes acceder a caracteres individuales contenidos dentro de la variable, usando corchetes [] y pasando el índice del caracter que quieres obtener:
* name[0] // 't' (índice 0)
* name[1] // 'e' (índice 1)

Puedes obtener una porción de la cadena indicándolo con esta sintaxis
* name[0:2] // obtiene: 'te'
* name[:2] // obtiene: 'te'
* name[2:] // obtiene: 'st'

Usando esto, puedes crear una copia de la cadena usando la siguiente sintaxis
* var newstring = name[:]

Puedes asignar la cadena a una nueva variable
* var first = "test"
* var second = first

Nota: Las cadenas son inmutables, no puedes actualizar el valor de una cadena.
Incluso si asignas un nuevo valor a la variable 'first' usando el operador de asignación (=) el valor de la variable 'second' siempre será "test"

* var first = "test"
* var second = "first"
* first = "another test"
* first // "another test"
* second //"test"

Las cadenas son tipos de referencia, lo que significa que si pasas una cadena como función, la referencia a esa cadena será copiada, no su valor. Dado que las cadenas son inmutables, in este caso no es una gran diferencia en la practica con pasar un tipo de dato entero, por ejemplo:

Puedes concatenar dos cadenas usando el operador aritmetico (+)
* var first = "first"
* var second = "second"
* var word = first + " " + second // "first second"

Go provee muchas utilidades de cadena en el paquete de cadenas (strings package)

Ejemplo:
Vamos a ver como importar un paquete en el ejemplo de "Hola mundo".

Así:

package main
import (
    "strings"
)

Y puedes usarlo, ejemplo: usamos la función HasPrefix() que sirve para observar si una cadena inicia con una subcadena especifica

package main
import (
    "strings"
)

func main() {
    string.HasPrefix("test","te")
}

Encontrar más métodos:
* https://pkg.go.dev/strings

Y hay más métodos como:
* strings.ToUpper()
* strings.ToLower()
* strings.HasSuffix()
* strings.Contains()
* strings.Count()
* strings.Join()
* strings.Split()
* strings.ReplaceAll()