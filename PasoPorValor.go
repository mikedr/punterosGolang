package main

import "fmt"

func main(){
  valorParaNegar := true
  fmt.Println("Valor antes del metodo: ", valorParaNegar)
  negar(valorParaNegar)
  fmt.Println("Valor despues del metodo: ", valorParaNegar)
}

func negar(valorParaNegar bool) bool {
  valorParaNegar = !valorParaNegar
  fmt.Println("Valor dentro del metodo: ", valorParaNegar)
  return valorParaNegar
}
