package main

import "fmt"

func removeElement(nums []int, val int) int {
	i, k := 0, 0
	j := len(nums) - 1
	for j >= i {
		if nums[i] == val && nums[j] == val {
			nums[j] = -1
			j--
			k++
		} else if nums[i] == val && nums[j] != val {
			nums[i] = nums[j]
			nums[j] = -1
			i++
			j--
			k++
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return len(nums) - k
}

func main() {
	nums := [4]int{3, 2, 2, 3}
	fmt.Println(removeElement(nums[:], 3))
	nums2 := [8]int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums2[:], 2))
	nums3 := [3]int{0, 0, 0}
	fmt.Println(removeElement(nums3[:], 2))
	nums4 := [1]int{1}
	fmt.Println(removeElement(nums4[:], 1))
}
