
# Cómo definir variables usando GO
Usamos la palabra reservada 'var'
* var edad = 20

Y las variables son definidas después de el o los paquetes principales (ejemplo: después de "fmt")

```
package main
import "fmt"
var edad = 20; <- aquí (es global)

func main() {
    fmt.Println("Hello, World!")
}

o dentro de la funcion main

package main
import "fmt"

func main() {
    var edad = 20; <- aquí (solo existe en la función)
    fmt.Println("Hello, World!")
}

```

Podemos declarar las variables sin inicialiarlas pero en ese caso debemos establecer el tipo de esta manera:
```

var edad int
var name string
var done bool

```

Cuando se conoce el valor de las variables usamos el operador ':=' (operador de variable corta)

```
age := 10
name := "Roger"
```

Adicional a esto se pueden usar como nombres de variable:
* Letras
* Digitos
* Guiones bajos

Puedes asignar un nuevo valor a la variale con el operador de asignación '='
```

var edad int
edad = 10
edad = 11
```

Se pueden declarar constantes así:
```
const edad = 10
```

Declarar multiples variables en una sola linea
```

var edad, name

e inicializarlas

var edad = 10, name = "Roger"
```
