# Punteros en Golang

## Paso por valor

En Go, al igual que muchos otros lenguajes, las variables se pasan "por valor" a las funciones. Esto significa que las funciones que son invocadas con parámetros reciben una copia de los mismos.

Esto significa que si la variablle pasada como parámetro a la funcion se modifica dentro del método, este cambio no se reflejará en el exterior ya que se pasó una copia de la variable, no la original.

Teniendo en cuenta esta pequeña [aplicación](https://github.com/mikedr/punterosGolang/blob/main/PasoPorValor.go), por ejemplo.  

## Paso por referencia
