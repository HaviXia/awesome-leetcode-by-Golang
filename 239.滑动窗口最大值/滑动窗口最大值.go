package _39_滑动窗口最大值

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

*/
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	for i, v := range nums {
		if i%k == 0 {
			prefixMax[i] = v
		} else {
			prefixMax[i] = max(prefixMax[i-1], v)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suffixMax[i] = nums[i]
		} else {
			suffixMax[i] = max(suffixMax[i+1], nums[i])
		}
	}

	ans := make([]int, n-k+1)
	for i := range ans {
		ans[i] = max(suffixMax[i], prefixMax[i+k-1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
