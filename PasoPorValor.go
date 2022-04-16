package main

import "fmt"

func main(){
  valor := true
  fmt.Println("Valor antes del metodo: ", valor)
  negar(valor)
  fmt.Println("Valor despues del metodo: ", valor)
}

func negar(valorParaNegar bool) bool {
  valorParaNegar = !valorParaNegar
  fmt.Println("Valor dentro del metodo: ", valorParaNegar)
  return valorParaNegar
}
