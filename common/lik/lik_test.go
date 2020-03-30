package lik

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	generate := Generate("[1,2,3,4,5]")
	fmt.Println(generate)
}
