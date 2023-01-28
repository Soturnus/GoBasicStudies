package main

import "fmt"

func main() {

	// len() Descobre o tamanho de uma string.
	fmt.Println(len("Riva"))
	//[] retorna o byte da posição solicitada, e não o caractere.
	fmt.Println("RivaJr"[1])
	// Contactenação
	fmt.Println("Riva " + "Jr")
}
