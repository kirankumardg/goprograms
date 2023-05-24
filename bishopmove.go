package main

import "fmt"
import "math"

func main() {

	var a = [3][3] int{{0,0,0},{0,0,0},{0,0,0}}

	fmt.Println(a[2][2]) 

	var i=2

        var j=2


	var c=0


	for k:=0;k<3;k++{
		for l:=0;l<3;l++{
		if (math.Abs(float64(i-k))<=1 || math.Abs(float64(j-l))<=1) && (k!=i&&l!=j) {
			c=c+1
		}
	}
	}

	fmt.Println(c)
		
}
