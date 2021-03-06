# Punteros en Golang

## Paso por valor

En Go, al igual que muchos otros lenguajes, las variables se pasan "por valor" a las funciones. Esto significa que las funciones que son invocadas con parámetros reciben una copia de los mismos.

Por lo tanto, si la variable pasada como parámetro a la funcion se modifica dentro del método, este cambio no se reflejará en el exterior ya que se pasó una copia de la variable, no la original.

Teniendo en cuenta esta [aplicación](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go), por ejemplo. [Aquí](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go#L7) se hace una impresión por pantalla del valor de la variable, el cual es `true`. Luego, [aquí](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go#L8) se invoca a la función con esa variable. [Este](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go#L12) método invierte su valor de verdad y lo imprime por pantalla, `false`. Finalmente, [aquí](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go#L9) se vuelve a imprimir por pantalla el valor de la variable, pero como esta acción se realiza fuera del contexto del método el valor sigue siendo `true`.

## Punteros

Es posible obtener la dirección de memoria de una variable usando el operando `&` (ampersand), el cual en Go se lee como "la dirección de". Por ejemplo, esta [aplicación](https://github.com/mikedr/punterosGolang/blob/main/ImprimeMem1.go) declara e inicializa una variable, luego imprime por pantalla su valor y "la dirección de" memoria de la misma.

Qué es una dirección de memoria? Es un referencia a una ubicación específica en la memoria, donde se pueden guardar variables de todo tipo:

|	Dirección de memoria	|	Valor que contiene |
|:---:|:---:|
|`0xc000040240`|	true	|
|`0xc000040248`|	25	|
|`0xc000040250`|	"River Plate"	|
|`0xc000040258`|	...	|
|`0xc000040260`|	...	|
|`0xc000040268`|	9.12	|

Los valores que representan una dirección de memoria son conocidos como **punteros**, ya que apuntan a la ubicación donde la variable se encuentra.

## Tipos punteros

El tipo de un puntero se escribe usando el síbolo `*` (asterísco), seguido del tipo de variable al que apunta. El tipo de un puntero a una variable `int`, por ejemplo, sería `*int`, lo cual se lee como "un puntero a un entero".

Entonces podemos declarar variables que contienen punteros. Un puntero puede solo contener punteros a un tipo de valor.

```
var miEntero int //declaración de un entero
var miPunteroAunEntero *int	//declaración de un puntero a un entero
miPunteroAunEntero = &miEntero	//asignación de un puntero a la dirección de memoria de la variable declarada
```

## Obteniendo o cambiando el valor de un puntero

Se puede obtener el valor de la variable a la que referencia un puntero escribiendo el operador `*` justo antes del puntero. 