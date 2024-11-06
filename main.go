package main



func main () {
    t := Task{
        Title: "dormir",
        Description: "c'est bien de dormir",
        Completed: false,
    }
    tasks[t.Title] = &t   // ou tasks["Dormir"] = t

    //log.printf("Notre tâche : %v", t)
    //t.display()

    displayAlltasks()

    addTask()
    t.Done()
    displayAlltasks()
}