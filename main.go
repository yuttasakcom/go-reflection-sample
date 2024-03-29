package main

import (
	"fmt"

	"github.com/yuttasakcom/go-reflection-sample/report"
)

type Person struct {
	name string `report:"ชื่อ,uppercase"`
	age  int    `report:"อายุ"`
}

type Employee struct {
	name string
	age  int
}

func main() {
	fmt.Println(report.Text(Person{name: "Yo", age: 36}))
	fmt.Println(report.Text(Employee{name: "Yea", age: 31}))
	fmt.Println(report.Text(struct {
		name string `report:",uppercase"`
		age  int
	}{name: "Sri", age: 32}))
}
