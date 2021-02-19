package main

import (
	"fmt"
)

func main() {
	var per string
	fmt.Println("Enter to permissions")
	fmt.Scan(&per)

	var cad string
	var fileType = per[0]
	switch fileType {
	case '-':
		cad = "Regular file with permission"
	case 'd':
		cad = "Directory with permission"
	case 'b':
		cad = "Special file with permission"
	case 'c':
		cad = "Special character file with permission"
	case 'l':
		cad = "Symbolic link with permission"
	case 'p':
		cad = "Pipe with permission"
	case 's':
		cad = "Socket with permission"
	}

	var perActiveClass int
	var perActiveTotal string

	/*
		Loop operation:
		per[1:4]
		per[4:7]
		per[7:9]
	*/
	for j := 0; j < 7; j += 3 {
		perActiveClass = 0
		for i := j; i < j+3; i++ {
			var char = per[i+1]

			if char == 114 { //r
				perActiveClass += 4
			}
			if char == 119 { //w
				perActiveClass += 2
			}
			if char == 120 { //x
				perActiveClass++
			}
		}
		perActiveTotal += fmt.Sprint(perActiveClass)
	}
	fmt.Print(cad, perActiveTotal)
}
