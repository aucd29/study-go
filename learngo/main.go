package main

import (
	"errors"
	"fmt"
	"hw/mydict"
	"hw/myrandstring"
	"log"
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

func (p *Hello) ChangeBalance(newBalance int) {
	p.balance = newBalance
}

func (p *Hello) GetBalance() int {
	return p.balance
}

func (p *Hello) ChangeDeveloper(changeDeveloper string) {
	p.devel = changeDeveloper
}

func (p *Hello) Withdraw(amount int) error {
	if p.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}

	p.balance -= amount
	return nil
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

	// WITHDRAW
	err := hello.Withdraw(100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello.GetBalance())

	// err2 := hello.Withdraw(2000)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// } else {
	// 	fmt.Println("ok")
	// }

	// DICTIONARY
	dictionary := mydict.Dictionary{
		"a": "b",
		"c": "d",
	}

	mapValue, err := dictionary.Find("a")
	fmt.Println(dictionary)
	if err == nil {
		fmt.Println(mapValue)
	}
	dictionary.Add("burke", "some value")

	mapValue2, err2 := dictionary.Find("burke")
	fmt.Println(dictionary)
	if err2 == nil {
		fmt.Println(mapValue2)
	}

	myrand := myrandstring.NewRandString()
	myrand.Init("burke")
	myrand.Hello()

	urls := []string{
		"https://wadiz.kr",
		"https://naver.com",
	}

	for _, url := range urls {
		fmt.Println("url : ", url)
	}

	// INTERFACE TEST
	var testInterface TestInterface

	testInterface = &TestStruct{"hello interface"}
	testInterface.TestMethod() // struct 에 해당하는 method 실행

	testInterface = TestFloat(111)
	testInterface.TestMethod() // float 에 해당하는 method 실행

}
