package main

import (
	"container/list"
	"fmt"
)
//stack and queue implementation
func main() {
	words := list.New()  //every time it increments the size by 1 on insert or adjust depends on operation


	for i:=0 ; i<5; i++{
		addElements(i , words)
	}
	//retrieve the data // Last in First Out - STACK implementation
	for i:=words.Back() ; i != nil ; i = i.Prev(){
		fmt.Println(i.Value) // fmt.println print the data and println print the address ?
	}
    println("________________________________")
	//retrieve the data // fifo queue
	for i:=words.Front() ; i != nil ; i = i.Next(){
		fmt.Println(i.Value) // fmt.println print the data and println print the address ?
	}

}

func addElements(data interface{}, words *list.List) {
	words.PushBack(data)
}

