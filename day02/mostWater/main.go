/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

*/

package main

import "fmt"

// 自己最开始想到的最笨的方法, 这个的时间复杂度是O(n^2)
func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			res := 0
			if height[i] > height[j] {
				res = (j - i) * height[j]
			} else {
				res = (j - i) * height[i]
			}
			if res > max {
				max = res
			}
		}
	}
	return max
}

// 这个方法的时间复杂度是O(N) ,左右边界i，j 向中间收敛
func maxArea2(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i != j; {
		res := 0
		if height[i] > height[j] {
			res = (j - i) * height[j]
			j--
		} else {
			res = (j - i) * height[i]
			i++
		}
		if res > max {
			max = res
		}
	}
	return max
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea2(height)
	fmt.Println(res)
}
