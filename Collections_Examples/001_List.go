package main

import (
	"container/list"
	"fmt"
)
//double linked list in go
func main() {
	words := list.New()

	println("the initial size of the list is ", words.Len()) // 0

	//push at the end
	for i:=0 ; i<6; i++{
		words.PushBack(i)  //0 , 1, 2, 3, 4, 5

	}

	//remove 3 from the list
	//initialize where to point the iterating pointer back or front
	for i:=words.Back() ; i != nil ; i=i.Prev(){
		if i.Value == 3{
			words.Remove(i)
		}
	}
	// insert after 2

	for i:=words.Back() ; i != nil ; i=i.Prev(){
		if i.Value == 2{
			words.InsertAfter(3, i)
		}
	}
	// interchanging 3 and 4
	for i:=words.Back() ; i != nil ; i=i.Prev(){
		if i.Value == 3 && i.Next() != nil{
			words.MoveAfter(i,i.Next())
		}
	}
	//interchanging 3 and 4
	for i:=words.Back() ; i != nil ; i=i.Prev(){
		if i.Value == 3 && i.Prev() != nil{
			words.MoveBefore(i,i.Prev())
		}
	}

	for i:=words.Back() ; i != nil ; i=i.Prev(){
		fmt.Println(i.Value)
	}






	
}
