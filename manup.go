package main

import "log"

var (
	articles   []string
	courseList = make(map[string]int)
)

func main() {

}

func displayCourse() {
	log.Printf("%v", articles)
	log.Printf("%v", courseList)
}