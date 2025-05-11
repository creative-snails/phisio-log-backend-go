package main

import (
	"fmt"
	"time"
)

type MyExample struct {
	Prop1 Prop1Type
}

type Prop1Type struct {
	regular  string
	template string
	function func(argument string) string
}

func NewMyExample() *MyExample {
	me := &MyExample{}

	me.Prop1.regular = "this is regular text"
	me.Prop1.template = fmt.Sprintf(`This is here with some dynamic value, for now it will the time: %s`, time.Now().Format("2006-01-02"))
	me.Prop1.function = func(argument string) string {
		return "Some logic was running and now we have this "+ argument
	}

	return me
}

func main() {
	myExample := NewMyExample()

	fmt.Println(myExample.Prop1.regular)
	fmt.Println(myExample.Prop1.template)
	fmt.Println(myExample.Prop1.function("Yooo"))
}