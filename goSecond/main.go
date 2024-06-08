package main

import (
	"errors"
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

	printMe(stringS)
	//	var charRune rune = 'h'
	fmt.Println(utf8.RuneCountInString(stringS))
	fmt.Println(multiply(25, 34))
	multi, add, err := multiply(26, 25)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Printf(err.Error())
	}
	fmt.Printf("the multiplication is %v and the addition is %v", multi, add)

	myVar := "hello bootyfull"

	var1, var2 := 24, "3455"
	fmt.Println(myVar, var1, var2)

	// arrays in go learning

	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0:3])
	fmt.Println(intArr[1])
	fmt.Println(intArr[2])
	fmt.Println(intArr)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func multiply(firstNum int, secondNum int) (int, int, error) {
	var err error
	if secondNum == 0 {
		err = errors.New("Second num cannot be zero")
		return 0, 0, err
	}
	var result int = firstNum * secondNum
	var addition int = firstNum + secondNum
	return result, addition, err
}
