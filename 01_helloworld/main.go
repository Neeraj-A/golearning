package main

import "fmt"

func main() {
	//fmt.Println("Hello all")
	fmt.Printf("Hello all")
	for i:= 1; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %U \n", i, i, i,i)
	}
}
