package main

import (
	"fmt"
	"lc/common/tri"
)

func main() {
	generate := tri.T{}.Generate("[1,null,2,3]")
	fmt.Println(generate)
}
