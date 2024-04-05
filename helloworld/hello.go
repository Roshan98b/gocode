package main

import (
	"fmt"
)

type engine struct {
	bhp int
	cc int
}

type car struct {
	company string
	price int
	engine
}

func main() {
	fmt.Println("Hello World!")

	// Struct
	mercedes := car{"Mercedes",300000, engine{300, 30000}}
	fmt.Println(mercedes)

	// Array
	array := [3]int{}
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array, len(array))

	// Slice
	slice := []int{}
	slice = append(slice, 1)
	slice = append(slice, array[1:]...)
	fmt.Println(slice, len(slice))

	// Pointer
	var num int = 200
	var p *int = &num
	var pp **int = &p
	fmt.Println("Value: ", num)
	fmt.Println("Value: ", p, " Pointer: ", *p, " Address: ", &p)
	fmt.Println("Value", pp, " Pointer: ", *pp, " Pointer of pointer: ", **pp, " Address: ", &pp)

}