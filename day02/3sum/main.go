/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// 复杂度O(n^2)
func threeSum(nums []int) [][]int {
	maps := make(map[int]int)
	resultMap := make(map[string]bool)
	result := make([][]int, 0)
	for i, v := range nums {
		for j := 0; j < i; j++ {
			// 这里判断的 v!=j 判断的是索引, 表示一个数只能使用一次
			if v, ok := maps[-v-nums[j]]; ok && v != j {
				entry := make([]int, 0)
				entry = append(entry, nums[i], nums[j], nums[v])
				sort.Ints(entry)
				// 这里提供了一种把切片转换为字符串的思路
				entryJson, err := json.Marshal(entry)
				if err != nil {
					panic(err)
				}
				entryStr := string(entryJson)
				if _, ok := resultMap[entryStr]; !ok {
					resultMap[entryStr] = true
					result = append(result, entry)
				}
			}
		}
		maps[v] = i
	}
	return result
}

// 重复写一遍
func threeSum2(nums []int) [][]int {
	maps := make(map[int]int)
	resultMap := make(map[string]bool)
	result := make([][]int, 0)
	for i, v := range nums {
		for j := 0; j < i; j++ {
			if v, ok := maps[-v-nums[j]]; ok && v != j {
				resultArray := []int{nums[i], nums[j], nums[v]}
				sort.Ints(resultArray)
				resultByte, err := json.Marshal(resultArray)
				if err != nil {
					panic(err)
				}
				resultStr := string(resultByte)
				if _, ok := resultMap[resultStr]; !ok {
					result = append(result, resultArray)
					resultMap[resultStr] = true
				}
			}
		}
		maps[v] = i
	}
	return result
}

func threeSum3(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for k, v := range nums {
		if v > 0 {
			return res
		}
		if k > 0 && v == nums[k-1] {
			continue
		}
		L := k + 1
		R := n - 1
		for L < R {
			if v+nums[L]+nums[R] == 0 {
				res = append(res, []int{v, nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L = L + 1
				}
				for L < R && nums[R] == nums[R-1] {
					R = R - 1
				}
				L = L + 1
				R = R - 1
			} else if v+nums[L]+nums[R] > 0 {
				R = R - 1
			} else {
				L = L + 1
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ret := threeSum2(nums)
	fmt.Println(ret)
}
