package main

import "fmt"

// m size of nums1
// n size of nums2
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := 0, 0; i < m; i++ {
		if nums1[i] > nums2[j] {
			aux := nums2[j]
			nums2[j] = nums1[i]
			nums1[i] = aux
		} else if nums1[i] == 0 && i >= n {
			nums1[i] = nums2[j]
			j++
		}
		if j == n {
			break
		}
	}
	fmt.Println(nums1)
}

func main() {
	arr1 := [6]int{1, 2, 3, 0, 0, 0}
	arr2 := [3]int{2, 5, 6}

	merge(arr1[:], 6, arr2[:], 3)
}
