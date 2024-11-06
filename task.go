package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
)

var (
	tasks = make(map[string]*Task)
)

type Task struct {
	Title string
	Description string
	Completed bool
}

// Pour clôturer une task
func (t *Task) Done() {
	t.Completed = true
}


func displayAlltasks() {
	for id, task := range tasks {
		fmt.Printf("\n-> Affichage de la tâche %v\n", id)
		task.Display()
	}
}

// crreer une nouvelle task avec saisie clavier et on l'ajoute dans la liste de tasks
func addTask() {
	fmt.Printf("\nSaisissez un titre: ")	
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)
	scan.Scan()
	titre := scan.Text()
	log.Printf("Titre: %v", titre)

	log.Printf("saisie une description: ")
	scan.Scan()
	desc := scan.Text()
	log.Printf("Description: %v", desc)

	t := Task{
		Title: titre,
		Description: desc,
		Completed: false,
	}
	// notre map de task stocke des pointers de task
	//on utilise donc le "&"
	tasks[titre] = &t
}

//cette fonction permet d'afficher une tâche, elle ne prend aucun argument et affiche sur la sortie standard
func (t Task) Display() {
	fmt.Printf("Titre: %v\n, Description: %v\n, Terminé: %v\n", t.Title, t.Description, t.Completed)
}


