package tri

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const Except = -math.MaxInt32

type Treer interface {
	Generate(str string) *TreeNode
}

type T struct {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//[1,null,2,3]
func (t T) Generate(str string) *TreeNode {
	res := t.strToArray(str)
	fmt.Printf("[Tree] StrToArray:%+v\n", res)
	return t.constructTree(res)
}

func (t T) constructTree(res []int64) *TreeNode {
	length := len(res)
	if length == 0 {
		return nil
	}
	var treeList = make([]*TreeNode, len(res))
	for i, v := range res {
		if v == Except {
			treeList[i] = nil
		} else {
			treeList[i] = &TreeNode{Val: int(v)}
		}
	}

	l_1 := 0
	l_2 := 1 // [l...r)
	l__1 := 0
	l__2 := 1 //[l...r) //使用滑动窗口思想
	for {
		var delNum int
		var i = l_1
		for ; i < l_2; i++ {
			if res[i] == Except {
				delNum++
			}
		}
		l__1 = l_1
		l__2 = l_2
		l_1 = i
		if (l_2-delNum)*2+1 > length {
			l_2 = length
		} else {
			l_2 = (l_2-delNum)*2 + 1
		}

		if l_1 >= length {
			break
		}

		for i := l_1; i < l_2; i++ { // i 之后
			for j := l__1; j < l__2 && i < l_2; j++ { //j之前
				if treeList[j] == nil {
					continue
				}
				//先左 后右
				treeList[j].Left = treeList[i]
				i++
				if i < l_2 {
					treeList[j].Right = treeList[i]
					i++
				}
			}
		}
	}
	return treeList[0]
}

func (t T) strToArray(str string) (res []int64) {
	str = strings.Trim(str, "[")
	str = strings.Trim(str, "]")
	split := strings.Split(str, ",")
	for _, v := range split {
		i, e := strconv.Atoi(v)
		if e != nil {
			i = Except
		}
		res = append(res, int64(i))
	}
	return
}

func GetTree() {

}
