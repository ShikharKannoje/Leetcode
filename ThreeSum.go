/*
3 Sum
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

//[1,1,1,1,3,1,2,4]

//j = 0
//m = 1

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	fmt.Println(ans)
}

//sorted [-4 -1 -1 0 1 2]
func threeSum(nums []int) [][]int {
	sol := make([][]int, 0)
	sort.Ints(nums)
	target := 0
	var i, j, k int
	for i = 0; i < len(nums)-1; {
		j = i + 1
		k = len(nums) - 1

		for j < k {

			if nums[j]+nums[k] < target-nums[i] {
				m := j + 1
				for nums[m] == nums[j] && m < len(nums)-1 {
					m++
				}
				j = m
			} else if nums[j]+nums[k] > target-nums[i] {
				n := k - 1
				for nums[n] == nums[k] && n > 0 {
					n--
				}
				k = n
			} else if nums[j]+nums[k] == target-nums[i] {
				temp := []int{nums[i], nums[j], nums[k]}
				sol = append(sol, temp)
				m := j + 1
				for nums[m] == nums[j] && m < len(nums)-1 {
					m++
				}
				j = m
				n := k - 1
				for nums[n] == nums[k] && n > 0 {
					n--
				}
				k = n
			}
		}
		l := i + 1
		for nums[l] == nums[i] && l < len(nums)-1 {
			l++
		}
		i = l
	}
	return sol
}
