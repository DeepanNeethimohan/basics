package main

import (
	"fmt"

	"github.com/DeepanNeethimohan/basics/dog/dog"
)

func main() {

	humanyears := 7

	dogyears := dog.Years(humanyears)

	fmt.Println(int(dogyears))
}
