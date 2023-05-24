package main

import "fmt"

func main() {
  	
    var str = "kirankumar"

    var length = len(str)

    var dict = map[string]int{}

    for i:=0;i<length;i++{
	
	var temp=string(str[i])
        val,ok:=dict[temp]
        if ok{
		dict[temp]=val+1	
	}else{
		dict[temp]=1
	}
	
    }
	fmt.Println(dict)

}

