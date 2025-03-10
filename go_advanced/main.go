package main

import (
	"fmt"
)

func PrintNum(x int) {
	fmt.Println(x)
}

func main() {
	var myChannel = make(chan string)
	var anotherChannel = make(chan string)

	go func() {
		myChannel <- "Hello"
	}()

	go func() {
		anotherChannel <- "World"
	}()

	/* var data = <-myChannel
	var data2 = <-anotherChannel

	fmt.Println(data2)
	fmt.Println(data) */

	select {
	case msgFromChannel := <-myChannel:
		fmt.Println(msgFromChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}
