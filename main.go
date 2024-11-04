package main

import (
	"fmt" 
	"bufio"
	"os"
)


func main() {
	//var task string
	var response string

	fmt.Print("Voulez-vous ajouter une tâche? (y/n): ")
	fmt.Scanln(&response)

	if response == "y" {
		fmt.Print("Veuillez entrer la description de la tâche: ")
		reader := bufio.NewReader(os.Stdin)
		task, _ := reader.ReadString('\n')
		fmt.Println("Tâche ajoutée:", task)
	} else {
		fmt.Println("Aucune tâche ajoutée.")
	}
	
}