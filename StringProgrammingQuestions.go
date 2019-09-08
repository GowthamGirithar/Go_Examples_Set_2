package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	StringContainDigits()
	IsStringNumber()
	NumberAndDigit()
	PalindromeCheck()
}

func PalindromeCheck() {
	var str1 = "Madam"
	str1=strings.ToLower(str1)
	for pos,_ := range str1{
		if str1[pos] != str1[utf8.RuneCountInString(str1)-1-pos] {
			println("Is not palindrome")
			break
		}
	}

}

func NumberAndDigit(){
	digit := rune('1')
	fmt.Println(unicode.IsDigit(digit))//true
	fmt.Println(unicode.IsNumber(digit))//true
	letter := rune('Ⅷ')
	fmt.Println(unicode.IsDigit(letter))//false
	fmt.Println(unicode.IsNumber(letter))//true
	other := rune('½')
	fmt.Println(unicode.IsDigit(other))//false
	fmt.Println(unicode.IsNumber(other))//true
}

func IsStringNumber() {
	var str ="Test"
	_, err := strconv.Atoi(str) // convert string to number
	if err == nil{
		println("is number")
	}
}

func StringContainDigits() {
	var str ="Test123"
	strings.ContainsAny(str, "s")// for sample
	for _, val := range str{ //return rune
		if unicode.IsDigit(val){ // for rune , it will unicode
			println("yes digit is there")
			break
		}


	}


}
