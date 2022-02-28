package main

import (
	"fmt"
	"time"
)

type TestIface interface {
	Greet() string
}

type Thing struct {}

func (i *Thing) Greet() string {
	return "Holla"
}


func main() {
	if i := Test(); i != nil {
		fmt.Println(i.Greet())
	}
}

func Test() TestIface {
	now := time.Now().Unix()

	if now % 2 == 0{
		return &Thing{}
	}

	return nil
}
