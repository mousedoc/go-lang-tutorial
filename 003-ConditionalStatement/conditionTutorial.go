package main

import "fmt"

func main() {
	ifTutorial()
	switchTutorial()
}

func ifTutorial() {
	i := 2

	if i == 2 {
		fmt.Println("i is 2")
	} else {
		fmt.Println("i is not 2")
	}

	j := 4

	if j == 3 {
		fmt.Print("j is three")
	} else if j == 4 {
		fmt.Println("j is four")
	} else {
		fmt.Println("i dont know")
	}

	if comp := i * 2; comp == 4 {
		fmt.Println("i x 2 = 4")
	}
}

func switchTutorial() {

	var name = "gimchi"

	switch name {
	case "kimchi":
		fmt.Println("this is kimchi")
	case "gimchi":
		fmt.Println("this is gimchi")
	default:
		fmt.Println("i dont know")
	}

	var i = 3

	switch i * 2 {
	case 6:
		fmt.Println("6")
	}

	switch comp := i * 2; comp - 1 {
	case 5:
		fmt.Println("5")
	}
}
