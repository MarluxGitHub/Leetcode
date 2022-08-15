/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			live := getNumber(board, i, j)
			next := 0
			if board[i][j] == 1 {
				if live >= 2 && live <= 3 {
					next = 1
				}
			} else {
				if live == 3 {
					next = 1
				}
			}
			if next > 0 {
				board[i][j] = board[i][j] | 2
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = board[i][j] >> 1
		}
	}
}

func getNumber(board [][]int, i, j int) int {
	live := 0
	directions := [][]int{
		[]int{-1, -1},
		[]int{-1, 0},
		[]int{-1, 1},
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, -1},
		[]int{1, 0},
		[]int{1, 1},
	}

	for _, direction := range directions {
		ii, jj := i+direction[0], j+direction[1]
		if ii < 0 || jj < 0 || ii >= len(board) || jj >= len(board[0]) {
			continue
		}

		if (board[ii][jj] & 1) == 1 {
			live++
		}
	}

	return live
}
// @lc code=end

