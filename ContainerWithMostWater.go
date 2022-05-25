/*
11. Container With Most Water
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	sol := maxArea(height)
	fmt.Println("Solution : ", sol)
}

func maxArea(height []int) int {
	var max int
	var i, j int
	i = 0
	j = len(height) - 1
	for i != j {
		a := height[i]
		b := height[j]
		lenght := min(a, b)
		width := j - i
		ar := area(lenght, width)
		if ar > max {
			max = ar
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
	return max
}

func area(a int, b int) int {
	return a * b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
