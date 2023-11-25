package main

import "fmt"

func removeElement(nums []int, val int) int {
	size := len(nums)
	idx_slice := make([]int, size)
	k := 0
	for j, i := 0, 0; i < size; i++ {
		if nums[i] == val {
			nums[i] = -1
			k++
		} else {
			idx_slice[j] = i
			j++
		}
	}
	idx_slice = idx_slice[:size-k]
	for j, i := 0, 0; i < size-k; i++ {
		if nums[i] == -1 {
			nums[i] = nums[idx_slice[j]]
			nums[idx_slice[j]] = -1
		}
		j++
	}
	fmt.Println(nums)
	return size - k
}

func main() {
	// nums := [4]int{3, 2, 2, 3}
	// fmt.Println(removeElement(nums[:], 3))
	// nums2 := [8]int{0, 1, 2, 2, 3, 0, 4, 2}
	// fmt.Println(removeElement(nums2[:], 2))
	nums3 := [3]int{0, 0, 0}
	fmt.Println(removeElement(nums3[:], 2))
}
