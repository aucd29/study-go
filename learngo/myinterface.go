package main

import "fmt"

type TestInterface interface {
	TestMethod()
}

type TestStruct struct {
	message string
}

// interface 상속을 특이하게 하네 -_ -;;;
func (t *TestStruct) TestMethod() {
	fmt.Println(t.message)
}

type TestFloat float64

// 타입에 대한 메소드 확장
func (f TestFloat) TestMethod() {
	fmt.Println(f)
}
