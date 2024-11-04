package main

import (
	"fmt"
	"log"
)

func add(a int, b int) int{
	return a + b
}

//var somme int. ou alors := pour déclarer et assigner en même temps une variable!

func main() {
	fmt.Println("hello world")
	somme := add(3, 4)
	fmt.Printf("La somme de 3 et 4 est : %v\n", somme)
	fmt.Printf("La multiplication de 3 et 4 est : %v\n", multiplication(3, 4))
	resultat, err := division(3, 0)
	if err != nil {
		log.Fatalf("alerte, division par zéro: %v", err)
	}
	log.Printf("La division de 3 et 4 est : %v\n", resultat)
	
}