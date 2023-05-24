package main

import "fmt"

func main() {

	var a = "initial"

	var length=len(a)

	fmt.Println(length)

	var b=""
	for i:=length-1;i>=0;i-- {
		b=b+string([]rune(a)[i])
	}

	fmt.Println(b)
}
