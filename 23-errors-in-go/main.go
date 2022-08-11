package main

import (
	"errors"
	"fmt"
)

	var numberIsNotEven = errors.New("The number is not even")
func IsEven(num int) error{
	if num  % 2 != 0 {
		return numberIsNotEven
	}
	return nil
}


func main() {

	fmt.Println("errors in go")
	IsEven(25)
	err:= IsEven(24)
	if err != nil {
		fmt.Println("Number is not even")
	}


	err = IsEven(25)
	if err != nil {
		// to handle multiple errors
		if err == numberIsNotEven{
			fmt.Println("numberIsNotEven error")
		}
		fmt.Println("number is not even")
		fmt.Println(err.Error())
	}
}
