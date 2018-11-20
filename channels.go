package main

import (
	"fmt"
)

func main() {

	c := putnumbers()

	printchannel(c)

	fmt.Println("Hello, playground")
}

func printchannel(c <-chan int) {

	for i := range c {
		fmt.Println(i)
	}
}

func putnumbers() <-chan int {

	c := make(chan int)

	go func() {

		for i := 0; i <= 99; i++ {

			c <- i

		}

		close(c)

	}()

	return c

}
