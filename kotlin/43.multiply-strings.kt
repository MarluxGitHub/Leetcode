/*
 * @lc app=leetcode id=43 lang=kotlin
 *
 * [43] Multiply Strings
 */

// @lc code=start
class Solution {
    fun multiply(num1: String, num2: String): String {
        if (num1 == "0" || num2 == "0") return "0"
        val n1 = num1.toCharArray().reversed().map { it - '0' }
        val n2 = num2.toCharArray().reversed().map { it - '0' }
        val res = IntArray(n1.size + n2.size)
        for (i in n1.indices) {
            for (j in n2.indices) {
                res[i + j] += n1[i] * n2[j]
            }
        }
        var carry = 0
        for (i in res.indices) {
            val sum = res[i] + carry
            carry = sum / 10
            res[i] = sum % 10
        }
        var i = res.size - 1
        while (i >= 0 && res[i] == 0) {
            i--
        }
        return if (i == -1) "0" else res.slice(0..i).reversed().joinToString("")
    }
}
// @lc code=end

