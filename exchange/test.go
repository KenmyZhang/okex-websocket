/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。



示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [0]
输出：0
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{}
	fmt.Println(getMaxContinueSum(nums))
}

func getMaxContinueSum(input []int) int {
	if len(input) == 0 {
		return 0
	}
	if len(input) == 1 {
		return input[0]
	}
	var maxSum int
	for i := 0; i < len(input); i++ {
		var preSum int
		for j := i; j < len(input); j++ {
			preSum = preSum + input[j]
			var tmpSum int
			for k := i; i < j; k++ {
				tmpSum = tmpSum + input[k]
			}
			if maxSum < tmpSum {
				maxSum = tmpSum
			}

		}
	}

	return maxSum
}
