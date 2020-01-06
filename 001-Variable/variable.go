package main

// Reserved words
// break        default      func         interface    select
// case         defer        go           map          struct
// chan         else         goto         package      switch
// const        fallthrough  if           range        type
// continue     for          import       return       var

// Variables
var intVariable int
var floatVariable float32 = .12 // .1f not work
var i, j, k int = 1, 2, 3

// Constants
const intConst int = 23
const stringConst string = "this is string"

// Const group should have 'COMMENT' as follows
// Card
const (
	Visa   = "Visa"
	Master = "MasterCard"
	Amex   = "American Express"
)

const (
	// First ordinal - or commend like this
	First = iota

	// Second ordinal
	Second

	// Third ordinal
	Third
)

func main() {
	println(intVariable)
	println(floatVariable)

	println(i, j, k)
	println(intConst)

	println(Visa)
	println(Master)
	println(Amex)

	println(First)
	println(Second)
	println(Third)

	// Short assignment statement (only able in method block)
	shortVar := 123
	println(shortVar)
}
