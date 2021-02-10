package array

import "testing"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}

	//处理 m=0 n != 0 的情况
	for i := 0; i < n; i++ {
		nums1[i] = nums2[i]
	}

}

func TestMerge(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	/*
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3
	*/
	merge(nums1, m, nums2, n)
	t.Log(nums1)
}
