package main

import (
	"errors"
	"fmt"
)

func IPanicEasily() {
	// defer is a function that is going to execute at the in of a function execution path
	defer func() {
		fmt.Println("1")
	}()
	panic("I panic easily")
}

func MyAwesomeFunction() (err error) {
	defer func() {
		// to recover the panic
		if r:= recover(); r!= nil {
			fmt.Println("recovered")
			err= errors.New("my  errooor panic")
		}
	    fmt.Println("2")

	}()
	IPanicEasily()
	return nil
}

func main() {
	defer func() {
		fmt.Println("3")
	}()
	fmt.Println("Panic! In the Go App")
	if err:= MyAwesomeFunction(); err != nil {
		fmt.Println("awww the funct returned an error")
		fmt.Println(err.Error())
	}
	fmt.Println("finished main function successfully")
}
