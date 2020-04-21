package main

import (
	"fmt"
	"sort"
)

func main() {
	sort := frequencySort("Aabb")
	fmt.Println(sort)
}

type C struct {
	x [128]int
	z [128]int
}

func (c C) Len() int {
	return 128
}

func (c *C) Less(i, j int) bool {
	return c.x[i] > c.x[j]
}

func (c *C) Swap(i, j int) {
	c.z[i], c.z[j] = c.z[j], c.z[i]
	c.x[i], c.x[j] = c.x[j], c.x[i]
}

func frequencySort(s string) string {
	c := &C{}
	for i := 0; i < 128; i++ {
		c.z[i] = i
	}
	for i := 0; i < len(s); i++ {
		c.x[s[i]] += 1
	}
	sort.Sort(c)
	var res string
	for i := 0; i < 128; i++ {
		jx := c.x[i]
		if jx == 0 {
			break
		}
		for j := 0; j < jx; j++ {
			res+=fmt.Sprintf("%c",c.z[i])
		}
	}
	return res
}
