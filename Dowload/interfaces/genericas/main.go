package main

import "fmt"

func generica(inter interface{}){//a interfacce generica é aceita 
	fmt.Println(inter)
}
func main() {
	generica(("lamlanla"))
	generica(1)
	generica(true)
}