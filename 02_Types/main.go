package main

import "fmt"

func main(){
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello" + " World")
	fmt.Println(true && false)
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(true || true)
	fmt.Println(!false)
	fmt.Println(32132 * 42452)
	fmt.Println((true && false) || (false && true) || !(false && false))
}
