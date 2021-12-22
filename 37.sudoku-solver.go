/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	blks, rows, cols := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for i := 0; i < 9; i++ {
		blks[i] = [9]bool{}
		rows[i] = [9]bool{}
		cols[i] = [9]bool{}
	}
	empty := [][]int{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				empty = append(empty, []int{i, j})
			} else {
				ch := int(board[i][j] - '1')
				rows[i][ch] = true
				cols[j][ch] = true
				blks[i/3*3+j/3][ch] = true
			}
		}
	}

	if helper(empty, 0, board, blks, rows, cols) {
		return
	}
}

func helper(empty [][]int, t int, board [][]byte, blks, rows, cols [9][9]bool) bool {
	if t == len(empty) {
		return true
	}
	i, j := empty[t][0], empty[t][1]
	for n := 1; n <= 9; n++ {
		if rows[i][n-1] || cols[j][n-1] || blks[i/3*3+j/3][n-1] {
			continue
		}
		rows[i][n-1] = true
		cols[j][n-1] = true
		blks[i/3*3+j/3][n-1] = true
		board[i][j] = byte(n + '0')
		if helper(empty, t+1, board, blks, rows, cols) {
			return true
		}
		rows[i][n-1] = false
		cols[j][n-1] = false
		blks[i/3*3+j/3][n-1] = false
		board[i][j] = '.'
	}
	return false
}
// @lc code=end

