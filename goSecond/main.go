package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("hello world")
	fmt.Println("1 + 1.0", 1+1.0)
	fmt.Println("1 / 2", 1.0/2)

	var intNum int = 9999999999999
	fmt.Println(intNum)
	var ans float64 = float64(intNum) + float64(24.5)
	fmt.Println(ans)
	var stringS string = `hello 
    what are you doing`

	fmt.Println(len(stringS))

	//	var charRune rune = 'h'
	fmt.Println(utf8.RuneCountInString(stringS))

	myVar := "hello bootyfull"

	var1, var2 := 24, "3455"
	fmt.Println(myVar, var1, var2)
}
