
# Operadores en GO
En GO se usan los siguientes operadores:
* =
* := 
* <

Y más...

Tenemos los operadores '=' y ':=', se usan para declarar e inicializar variables
```
var a = 1
b := 1
```

Tenemos dos operadores de comparación '==' y '!=' que toman 2 argumentos y retornan un booleano
```
var num = 1
num == 1 // true
num != 1 // false
```

También '<', '<=','>', '>='
```
var num = 1
num > 1 // false
num >= 1 // true
num < 1 // false
num <= // true
```

Operadores aritmenticos como:
```
1 + 1 // 2
1 - 1 // 0
1 * 2 // 2
2 / 2 // 1
2 % 2 // 0
```

El operador de suma puede concatenar strings
```
"a" + "b" //ab
```

También existen los operadores unarios (aumentan o decrementan una unidad)
```
var num =  1
num++ // 2
num-- // 1
```

Operadores booleanos como:
```
true && true  //true
true && false //false
true || false //true
false || false //false
!true  //false
!false //true
```