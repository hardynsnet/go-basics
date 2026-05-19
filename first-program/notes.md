# Primer programa con GO

El objetivo es imprimir por pantalla la frase: 'Hello World' y se hace teniendo en cuenta la sintaxis:

```
package main
import "fmt"

func main() {
    fmt.Println("Hello World");
}
```

## ¿Qué significa cada línea?

* package main: los programas de Go son organizados en paquetes, cada archivo '.go' primero declara a que paquete pertenece siendo así, un paquete puede estar compuesto por múltiples archivos o solo un archivo, un programa de Go puede contener múltiples paquetes.
* main: es el punto central del programa y lo identifica como un ejecutable
import: usamos esta palabra clave para importar un paquete, en este caso 'fmt' (es un paquete integrado de Go que contiene funciones de entrada y salida)
* func main () {}: es la función inicial del programa, y dentro de ella inicia todo. ejemplo: fmt.Println("Hello, World!")