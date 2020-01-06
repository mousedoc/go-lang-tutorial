package main

import "fmt"

// Difference between fmt.Println() and println()
// URL : https://stackoverflow.com/questions/14680255/whats-the-difference-between-fmt-println-and-println-in-go

func main() {
	boolTutorial()
	stringTutorial()
	intTutorial()
	floatTutorial()
	typeConversionTutorial()
}

func boolTutorial() {
	var flag = true
	fmt.Println(flag)

	flag = false
	fmt.Println(flag)
}

func stringTutorial() {
	// this is ` not a ---> '
	var raw = `This is \nstring`
	var interpreted = "This is \nstring"
	var rawExtended = `SLAM\n
SLAM
		S L A M`

	interpreted = "asdfasdf"

	fmt.Println(raw)
	fmt.Println(interpreted)
	fmt.Println(rawExtended)
}

func intTutorial() {
	// int	int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr

	intVar := -112318903238
	var uintVar uint = 1231232

	fmt.Println(intVar)
	fmt.Println(uintVar)
}

func floatTutorial() {
	// float32 float64 complex64 complex128
	floatVar := 112318903238.1231235789

	fmt.Println(floatVar)
}

func typeConversionTutorial() {
	floatVar := 123.42
	intVar := int(floatVar)
	uintVar := uint(intVar)
	stringVar := fmt.Sprintf("%f", floatVar)
	bytes := []byte(stringVar)

	fmt.Println(floatVar)
	fmt.Println(intVar)
	fmt.Println(uintVar)
	fmt.Println(stringVar)
	fmt.Println(bytes)
}
