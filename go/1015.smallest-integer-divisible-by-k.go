/*
 * @lc app=leetcode id=1015 lang=golang
 *
 * [1015] Smallest Integer Divisible by K
 */

// @lc code=start
func smallestRepunitDivByK(K int) int {
	if K % 2 == 0 || K % 5 == 0 { return -1 }
    cur, res := 0, 1
    for {
        cur = (cur * 10 + 1) % K
        if cur == 0 { return res }
        res += 1
    }
}
// @lc code=end

asdasdasdasdasdalsdkalsökdalösdklöaskdlöaskdalöskdalöskdlöasdgfjpowjeropewropdfsölafkadsalsödkfdlsäjgakäspd