/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
    	dirs := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		} else if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, dir)
		}
	}
	return "/" + strings.Join(stack, "/")
}
// @lc code=end

