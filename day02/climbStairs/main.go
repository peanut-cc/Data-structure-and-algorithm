/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

这个问题的分析过程
1层楼梯，你只有1种方法
2层楼梯，其实把问题想成，你从1层楼梯到2层，你有几种方法，只有一种，所以这里是2种方法
3层楼梯，其实是从第2层到第3层有几种方法，然后加上第二层的方法数
其实这个和fn = f(n -1) + f(n-2)

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
