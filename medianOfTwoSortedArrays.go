/*
4. Median of Two Sorted Arrays
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	sol := findMedianSortedArrays(nums1, nums2)
	fmt.Println("Solution : ", sol)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sol []int
	var i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			sol = append(sol, nums1[i])
			i++
		} else {
			sol = append(sol, nums2[j])
			j++
		}
	}
	if j < len(nums2) {
		for j < len(nums2) {
			sol = append(sol, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		for i < len(nums1) {
			sol = append(sol, nums1[i])
			i++
		}
	}

	ans := median(sol)
	return ans
}

func median(nums []int) float64 {
	var sol float64
	if len(nums)%2 == 0 {
		n1 := len(nums) / 2
		n2 := n1 + 1
		fmt.Println(n1, n2)
		sol = mean(nums[n1-1], nums[n2-1])
	} else {
		idx := (len(nums) + 1) / 2

		sol = float64(nums[idx-1])
	}
	return sol
}

func mean(a int, b int) float64 {
	a1 := float64(a)
	b1 := float64(b)
	return (a1 + b1) / 2
}
