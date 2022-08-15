/*
 * @lc app=leetcode id=329 lang=javascript
 *
 * [329] Longest Increasing Path in a Matrix
 */

// @lc code=start
/**
 * @param {number[][]} matrix
 * @return {number}
 */
 const longestIncreasingPath = (matrix) => {
    if (matrix == null || matrix.length === 0) return 0;
    const h = matrix.length;
    const w = matrix[0].length;
    const dirs = [[-1, 0], [0, 1], [1, 0], [0, -1]];
    const cache = [...Array(h)].map(() => Array(w).fill(null));

    const go = (x, y) => {
      if (cache[x][y] != null) return cache[x][y];
      let max = 1;
      for (const [dx, dy] of dirs) {
        const i = x + dx;
        const j = y + dy;
        if (i >= 0 && i < h && j >= 0 && j < w && matrix[i][j] > matrix[x][y]) {
          const len = go(i, j) + 1;
          max = Math.max(max, len);
        }
      }
      cache[x][y] = max;
      return max;
    };

    let max = 1;
    for (let i = 0; i < h; i++) {
      for (let j = 0; j < w; j++) {
        const len = go(i, j);
        max = Math.max(max, len);
      }
    }
    return max;
  };
// @lc code=end

