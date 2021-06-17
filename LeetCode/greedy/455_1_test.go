package greedy

import (
	"sort"
	"testing"
)

//g 胃口 s size
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gInd, sInd := 0, 0
	var child int

	for gInd < len(g) && sInd < len(s) {
		if g[gInd] <= s[sInd] { //满足
			child++
			gInd++ //胃口+1
		}
		sInd++
	}
	return gInd

}

func TestFindContentChildren(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 1}

	t.Log(findContentChildren(g, s))

}
