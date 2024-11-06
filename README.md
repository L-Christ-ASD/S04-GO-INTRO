# gotest
## 1. Créeez un fichier main.go
* *projet/main.go*

## 2. Implémmentez le programe de base:

-Créer une variable pour stockez la tâche
-Demandez à l'utilisateur s'il veut ajouter une tâche (y/n)
-Si oui, Demandez la déscription de la tâche
-Stocker la tâche dans la variable
-Afficher la tâche enregistrée

## 3. structure du programme:

```go

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
```
