package queue

import "testing"

type MovingAverage struct {
	data []int
	size int
}

func Constructor(size int) MovingAverage {
	init_data := make([]int, 0)
	return MovingAverage{
		data: init_data,
		size: size,
	}
}

func sum(arrs []int) int {
	res := 0
	for _, v := range arrs {
		res += v
	}
	return res
}

func (this *MovingAverage) Next(val int) float64 {
	clen := len(this.data)
	if clen < this.size {
		this.data = append(this.data, val)
	} else { //每次size 超过赋值
		this.data = this.data[1:]
		this.data = append(this.data, val)
	}
	clen = len(this.data)
	res := float64(sum(this.data)) / float64(clen)
	return res
}

func TestConsTructor(t *testing.T) {
	c := Constructor(3)
	t.Log(c.Next(3))
}
