/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

package main

import "fmt"

func main() {
	array := []int{2, 7, 11, 15}
	target := 9
	sol := twoSum(array, target)
	fmt.Println(sol)
}

func twoSum(nums []int, target int) []int {
	sol := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				sol = append(sol, i)
				sol = append(sol, j)
				return sol
			}
		}
	}
	return sol
}
