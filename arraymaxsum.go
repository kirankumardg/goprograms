package main

import "fmt"

func main() {
    var array = [] int { 1,5,21,13,6,4,8,2,9,10}

    var length = len(array)

    var fsuma =[10]int{}
    var bsuma = [10]int{}

    fsuma[0]=array[0]

    bsuma[0]= array[length-1]
   for i:=1;i<length;i++ {
       fsuma[i]= fsuma[i-1]+array[i]
       bsuma[i]=bsuma[i-1]+array[length-i-1]
       
 
   }

  fmt.Println(fsuma)
  fmt.Println(bsuma)

}
