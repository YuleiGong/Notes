package greedy

import (
	"sort"
	"testing"
)

//g 胃口 s size
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var count int
	var sIndex int
	var gIndex int

	for gIndex < len(g) && sIndex < len(s) {
		if g[gIndex] <= s[sIndex] {
			count++
			gIndex++
		}
		sIndex++
	}

	return count

}

func TestFindContentChildren(t *testing.T) {
	g := []int{1, 2}    //胃口
	s := []int{1, 2, 3} //size

	t.Log(findContentChildren(g, s))

}
