/*
 * @lc app=leetcode id=952 lang=kotlin
 *
 * [952] Largest Component Size by Common Factor
 */

// @lc code=start
class Solution {
    fun largestComponentSize(A: IntArray): Int {
        val count = A.toList().groupingBy { it }.eachCount().toMutableMap()
        val left = mutableMapOf<Int, MutableSet<Int>>()
        val right = mutableMapOf<Int, MutableSet<Int>>()
        val prime = BooleanArray(100001, { _ -> true })
        for (p in 2 until 100001) {
            if (prime[p]) {
                for (d in IntProgression.fromClosedRange(p, 100000, p)) {
                    if (d > p) prime[d] = false
                    if (count.containsKey(d)) {
                        left.computeIfAbsent(p, { _ -> hashSetOf() }).add(d)
                        right.computeIfAbsent(d, { _ -> hashSetOf() }).add(p)
                    }
                }
            }
        }

        var max = 0
        while (left.isNotEmpty()) {
            val queue = mutableSetOf<Int>()
            var sum = 0
            queue.add(left.keys.first())
            while (queue.isNotEmpty()) {
                val l = queue.first().also { queue.remove(it) }
                val rs = left.remove(l) ?: emptyList<Int>()
                for (r in rs) {
                    sum += count.remove(r) ?: 0
                    queue.addAll((right.remove(r) ?: mutableSetOf()).toList())
                }
            }
            max = kotlin.math.max(max, sum)
        }

        return max
    }
}
// @lc code=end

