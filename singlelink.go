package main

import "fmt"

type Node struct {
	Value int
        Next *Node

}

type Linkedlist struct {

	Size int
	Head *Node
	Tail *Node
}



func main() {
   n:=&Node{ Value:5}
   l:=&Linkedlist {Size:1,Head:n,Tail:n }
   fmt.Println(l.Size)	

   n1:=&Node{ Value:10}
   l.Size=2
   l.Head.Next=n1
   l.Tail=n1
   fmt.Println(l.Tail.Value)  

}

