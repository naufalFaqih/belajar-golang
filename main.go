package main

import (
	"belajar-golang-level-akses/library"
	"fmt"
)

func main(){
	var s1 = library.Student{"ethan", 21}
	fmt.Println("name", s1.Name)
	fmt.Println("name", s1.Grade)

}