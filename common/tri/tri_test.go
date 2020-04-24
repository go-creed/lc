package tri

import (
	"fmt"
	"testing"
)

func TestTreeNode_Generate(t *testing.T) {
	generate := T{}.Generate("[1,null,2,3]")
	fmt.Println(generate)
}
