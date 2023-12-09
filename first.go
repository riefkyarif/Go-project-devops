package main

import (
	"fmt"
	//"myFunctions"
)

func main() {
	fmt.Println("hello world!")
	looping()
	arrayyys()
	//myFunctions.FuncA()
	//myFunctions.FuncB()
}

func looping() {
	for i := 1; i < 11; i++ {
		fmt.Print(i, " ")

	}
	fmt.Println(" ")
}

func arrayyys() {
	//assign normal array berjumlah 9. dengan seluru angka diassign di awal.
	a := [9]int{1, 2, 9, 4, 5, 8, 0, 2, 1}
	fmt.Println(a)

	// buat array berjumlah 10 dengan initial "0". dengan assign 7 di index 1 dan 8 di index 5.
	ab := [10]int{}
	b := [10]int{1: 7, 5: 8}
	fmt.Println("create an initial array filled with zero, and add a value at a specified index:", b)
	fmt.Println("creaye an initial array filled with zero:", ab)

	//for printing individual index in array, use below, Kind of similar to that in Python.
	fmt.Println(a[5], b[7])

	//if you wanna create a sting-containing array, use the code below:
	c := [3]string{"apa", "kabar", "kamu?"}
	fmt.Println(c, "|", c[2])

	//if you want to change an individual index in an array, you can use the code below:
	c[2] = "LO Lo semua?"
	fmt.Println(c)
}
