/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

 */

package main

import "fmt"

// 这是最笨的方法,复杂度未O(N^2)
func twoSum(nums []int, target int) []int {
	ret := make([]int,2)
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i] + nums[j] == target {
				ret[0] = i
				ret[1] = j
			}
		}
	}
	return ret
}
// 复杂度O(N)
func towSum2(nums []int, target int) []int {
	// map的key是nums的值，value是nums的索引
	m:= map[int]int{}
	for i, v := range nums {
		if k, ok := m[target -v];ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil

}

func main() {
	nums := []int {2, 7, 11, 15}
	ret := twoSum(nums,9)
	fmt.Println(ret)
}
