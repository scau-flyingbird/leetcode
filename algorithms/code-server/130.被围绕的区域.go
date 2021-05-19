/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
var flag [][]bool
var m, n int
var sx, sy []int

func dfs(x int, y int, board [][]byte) {
    flag[x][y] = true
    for i := 0 ;i < 4; i++ {
        tx := x + sx[i]
        ty := y + sy[i]
        if tx < 0 || tx >= m || ty < 0 || ty >= n || board[tx][ty] == 'X' || flag[tx][ty] {
            continue
        }
        dfs(tx, ty, board)
    }
}

func solve(board [][]byte)  {
    m = len(board)
    if m == 0 {
        return
    }
    n = len(board[0])
    sx = []int{-1, 1, 0, 0}
    sy = []int{0, 0, -1, 1}

    flag = make([][]bool, m)
    for i := 0 ;i < m ;i++ {
        flag[i] = make([]bool, n)
    }

    for i := 0; i < m; i++ {
        for j := 0 ; j < n; j++ {
            if (i == 0 || i == m - 1 || j == 0 || j == n - 1) && board[i][j] == 'O' && !flag[i][j]{
                dfs(i, j, board)
            }
        }
    }

    for i := 0 ; i < m ; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' && !flag[i][j] {
                board[i][j]='X'
            }
        }
    }
}
// @lc code=end

