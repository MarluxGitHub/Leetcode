package leetcode

/*
 * @lc app=leetcode id=393 lang=golang
 *
 * [393] UTF-8 Validation
 */

// @lc code=start
func validUtf8(data []int) bool {
    var n int
    for _, b := range data {
        if n == 0 {
            if b>>5 == 0b110 {
                n = 1
            } else if b>>4 == 0b1110 {
                n = 2
            } else if b>>3 == 0b11110 {
                n = 3
            } else if b>>7 != 0 {
                return false
            }
        } else {
            if b>>6 != 0b10 {
                return false
            }
            n--
        }
    }
    return n == 0
}
// @lc code=end

