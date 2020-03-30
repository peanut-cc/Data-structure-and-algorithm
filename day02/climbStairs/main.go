/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
*/

package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	_, f2, f3 := 1, 2, 3
	for i := 4; i <= n; i++ {
		f2, f3 = f3, f2+f3
	}
	return f3
}

func main() {
	res := climbStairs(4)
	fmt.Println(res)
}
