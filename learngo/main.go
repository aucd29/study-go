package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (lenName int, upperName string) {
	defer fmt.Println("done")
	lenName = len(name)
	upperName = strings.ToUpper(name)

	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
		fmt.Println(num)
	}

	return sum
}

func can_I_drink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}

type Hello struct {
	devel   string
	balance int
}

func NewHello(developer string) *Hello {
	return &Hello{devel: developer, balance: 0}
}

func (hello *Hello) ChangeBalance(newBalance int) {
	hello.balance = newBalance
}

func (hello *Hello) GetBalance() int {
	return hello.balance
}

func main() {
	// HELLOWORLD
	fmt.Println("hello world")
	const name string = "burke"
	fmt.Println(name)

	// LENGTH, AND UPPER CASE
	totalLength, upperName := lenAndUpper(name)
	fmt.Println(totalLength, upperName)

	// MULTIPLE STRING
	repeatMe("a", "b", "c", "d", "e")

	// DIFFER
	totalLength2, upperName2 := lenAndUpper2(name)
	fmt.Println(totalLength2, upperName2)

	// 	SUPER ADD
	resultAdd := superAdd(1, 2, 3, 4)
	fmt.Println(resultAdd)

	// DRINK
	fmt.Println(can_I_drink(19))
	fmt.Println(can_I_drink(18))
	fmt.Println(can_I_drink(17))

	// METHOD
	hello := NewHello("bk")
	hello.ChangeBalance(2000) // SET 2000
	fmt.Println(hello.GetBalance())

}
