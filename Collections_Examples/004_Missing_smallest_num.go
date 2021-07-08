package main

import "fmt"

func main() {
	//sorted array. find smallest missing number
  //no duplicate and always start with 0
	arr:=[]int{0,1,8,9,10}
	//linear
	for i, v := range arr{
		if i !=v{
			fmt.Println("Missing number is ",i)
			break
		}
	}
	//binary search
	data:=findSmallestNumber(arr,0,len(arr)-1)
	fmt.Println("Missing number via recurssion is ",data)


}

func findSmallestNumber(arr []int, start , end int) int{
	if len(arr)==0{
		return 0
	}

	if arr[start] != start{
		return start
	}

	midLength:= (start+end)/2

	if arr[midLength] != midLength{
		return findSmallestNumber(arr,start,midLength)
	}

	return findSmallestNumber(arr,midLength+1, end)
}

