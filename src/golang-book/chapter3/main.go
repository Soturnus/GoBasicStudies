package main

import "fmt"

func main() {

	part1 := "3.1. Variaveis \n"
	fmt.Println(part1)

	var nome string = "RivaJr \n"
	fmt.Println(nome)

	// atribuição de concatenação
	var x string = "first"
	fmt.Println(x)
	x = x + " second \n"
	fmt.Println(x)

	var firstName string = "Riva"
	var lastName string = "Júnior"
	fmt.Println(firstName == lastName)

	fullName := "Riva Júnior \n"
	fmt.Println(fullName)

	dogsName := "Max \n"
	fmt.Println("My dog's name is: ", dogsName)

	part2 := "3.2. Constantes \n"
	fmt.Println(part2)

	constantes()
}

func constantes() {

	fmt.Println("Teste")
}
