package myrandstring

import (
	"fmt"
	"math/rand"
	"time"
)

type RandString struct {
	_name string
}

func NewRandString() *RandString {
	return &RandString{_name: ""}
}

func (p *RandString) Init(name string) {
	rand.Seed(time.Now().UnixNano())
	p._name = name
}

func (p *RandString) randomString() string {
	formats := []string{
		"Hi %v Welcome!",
		"Great to see you, %v!",
		"Hail, %v Well met!",
		"안녕 %v 환영해!",
	}

	return formats[rand.Intn(len(formats))]
}

func (p *RandString) Hello() {
	message := fmt.Sprintf(p.randomString(), p._name)
	fmt.Println("==========")
	fmt.Println(message)
	fmt.Println("==========")
}
