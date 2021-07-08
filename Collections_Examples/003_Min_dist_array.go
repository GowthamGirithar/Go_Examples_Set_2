package main

import (
	"fmt"
	"math"
)

func main() {
	//Find min distance in an unsorted array
	arr:=[]int{1,30,10,10,20,4,9,9}
	x:=1
	y:=20
	min:=math.MaxInt64

	for i:=0; i<len(arr) ;i++{
		for j:=i+1; j< len(arr); j++{
			if (arr[i]==x && arr[j]==y) || (arr[i]==y && arr[j]==x){
				if min > j-i{
					min= j-i
				}
			}
		}
	}

	fmt.Println(min)
	fmt.Println(math.Abs(-8))// to avoid negative number //8


}

