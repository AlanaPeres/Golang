package main //conversão é quando tomamos um valor de um certo tipo e o transformamos em outro tipo.

import (
 "fmt"
)

var a int

type carro int 

var b carro

func main() {
  a = 42
  b = 43
  fmt.Println(a)
  fmt.Printf("%T\n", a)
  fmt.Println(b)
  fmt.Printf("%T\n", b)
}

/*
42
int       
43        
main.carro
*/