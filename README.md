# Punteros en Golang

## Paso por valor

En Go, al igual que muchos otros lenguajes, las variables se pasan "por valor" a las funciones. Esto significa que las funciones que son invocadas con parámetros reciben una copia de los mismos.

Por lo tanto, si la variable pasada como parámetro a la funcion se modifica dentro del método, este cambio no se reflejará en el exterior ya que se pasó una copia de la variable, no la original.

Teniendo en cuenta esta [aplicación](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go), por ejemplo. [Aquí](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go#L7) se hace una impresión por pantalla del valor de la variable, el cual es `true`. Luego, [aquí](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go#L8) se invoca a la función con esa variable. [Este](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go#L12) método invierte su valor de verdad y lo imprime por pantalla, `false`. Finalmente, [aquí](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go#L9) se vuelve a imprimir por pantalla el valor de la variable, pero como esta acción se realiza fuera del contexto del método el valor sigue siendo `true`.

## Punteros

Es posible obtener la dirección de memoria de una variable usando el operado & (ampersand), el cual en Go se lee como "la dirección de". Por ejemplo, esta [aplicación](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go) declara e inicializa una variable, luego imprime por pantalla su valor y "la dirección de" memoria de la misma.

Qué es una dirección de memoria? Es un referencia a una ubicación específica en la memoria, donde se pueden guardar variables de todo tipo:

|---|:---:|
|`0xc000040240`|	true	|
|`0xc000040248`|	25	|
|`0xc000040250`|	"River Plate"	|
|`0xc000040258`|	...	|
|`0xc000040260`|	...	|
|`0xc000040268`|	9.12	|