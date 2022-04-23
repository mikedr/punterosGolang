package main

import "fmt"

func main(){
  var num int = 25 //se declara y asigna una variable del tipo int
  var puntNum *int //se declara un variable tipo puntero a un entero
  puntNum = &num //se asigna el puntero hacia la dirección de memoria de puntNum
  fmt.Println("El puntero apunta a la dirección: ",puntNum) //se imprime el valor del puntero (dirección a la que apunta)
  fmt.Println("La dirección a la que apunta contiene el valor: ",*puntNum) //se imprime el valor que contiene la dirección de memoria a la que apunta el puntero)
}
