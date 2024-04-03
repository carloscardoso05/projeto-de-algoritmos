package main

import (
	"fmt"
	"main/numero_complexo"
)

func main() {
	n1 := numero_complexo.Criar(10, 2)
	n2 := numero_complexo.Criar(5, 0)
	
	n1.Imprimir()
	fmt.Printf("N1 é real: %t\n", n1.EhReal())
	n2.Imprimir()
	fmt.Printf("N2 é real: %t\n", n2.EhReal())
	
	n2 = numero_complexo.Somar(n1, n2)
	n2.Imprimir()
	fmt.Printf("N2 é real: %t\n", n2.EhReal())
}