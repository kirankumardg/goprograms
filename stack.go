package main

import "fmt"

type Stack []string

func (s *Stack) push ( str string ){
   *s=append(*s,str)
}

func (s *Stack) pop() (string, bool){
   if len(*s)==0{
	fmt.Println("stack is empty no element to pop")
	return "",false
   }else{
         index:=len(*s)-1
         element:=(*s)[index]
         *s=(*s)[:index]
         return element, true
	
	}



}

func main() {
	var stack Stack

	stack.push("Kiran")
	stack.push("arun")

	fmt.Println(stack.pop())

}

