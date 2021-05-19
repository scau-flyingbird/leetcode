/*
 * @lc app=leetcode.cn id=1799 lang=golang
 *
 * [1799] N 次操作后的最大分数和
 */
 // @lc code=start
 var g [][]int
 var n int
 var maxstate int
 var f []int
 
 func gcd(a, b int) int {
	 if a == b {
		 return a
	 }
	 evena, evenb := a & 1, b & 1
	 if evena == 0 && evenb == 0 {
		 return gcd(a >> 1, b >> 1) << 1
	 }else if evena == 1 && evenb == 0 {
		 return gcd(a, b >> 1)
	 }else if evena == 0 && evenb == 1 {
		 return gcd(a >> 1, b)
	 }else {
		 if b > a {
			 a, b = b, a
		 }
		 return gcd(a-b, b)
	 }
 }
 
 func max(a, b int) int {
	 if a > b {
		 return a
	 }else {
		 return b
	 }
 }
 
 func dfs(t int, state int) int {
	 if f[state] != -1 {
		 return f[state]
	 }
	 for i := 0; i < n; i++ {
		 for j := i + 1; j < n; j++ {
			 if state & (1 << i) == 0 || state & (1 << j) == 0 {
				 continue
			 }
			 newstate := state - (1 << i) - (1 << j)
			 f[state] = max(f[state], dfs(t + 1, newstate) + t * g[i][j])
		 }
	 }
	 return f[state]
 }
 
 func maxScore(nums []int) int {
	 n = len(nums)
	 g = make([][]int, n)
	 maxstate = 1 << n
	 f = make([]int, maxstate)
 
	 for i := 0 ; i < len(nums); i++{
		 g[i] = make([]int, n)
		 for j := i + 1; j < len(nums);j++ {
			 g[i][j] = gcd(nums[i], nums[j])
		 }
	 }
 
	 for i := 0 ;i < maxstate; i++{
		 f[i] = -1
	 }
	 f[0] = 0
 
	 return dfs(1, maxstate - 1)
 }
// @lc code=end

