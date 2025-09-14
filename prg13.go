package main

func twoSum(nums []int, target int) []int {
	retArr := make([]int, 2)

	for i, v1 := range nums {
		for j, v2 := range nums {
			if j != i && v1+v2 == target {
				retArr[0] = i
				retArr[1] = j
				return retArr
			}
		}
	}
	return retArr
}
