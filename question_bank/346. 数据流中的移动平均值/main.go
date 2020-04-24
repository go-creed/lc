package main

import "fmt"

func main() {

	constructor := Constructor(3)

	fmt.Println(constructor.Next(1))
	fmt.Println(constructor.Next(10))
	fmt.Println(constructor.Next(3))
	fmt.Println(constructor.Next(5))
}

type MovingAverage struct {
	data []int
	size int
	left int
	sum  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	avarage := MovingAverage{}
	avarage.size = size
	avarage.left = 0
	avarage.data = append(avarage.data, 0)
	return avarage
}

func (this *MovingAverage) Len() int {
	return len(this.data) - 1
}

func (this *MovingAverage) Next(val int) float64 {
	this.data = append(this.data, val)

	if this.Len() > this.size {
		this.left++
		this.sum -= this.data[this.left]
	}
	this.sum += val
	if this.Len() >= this.size {
		return float64(this.sum) / float64(this.size)
	}
	return float64(this.sum) / float64(this.Len())
}
