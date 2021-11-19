/*
 * @lc app=leetcode id=1178 lang=kotlin
 *
 * [1178] Number of Valid Words for Each Puzzle
 */

// @lc code=start
class Solution {
    fun findNumOfValidWords(words: Array<String>, puzzles: Array<String>): List<Int> {
        val wordMaskMap:MutableMap<Int,Int> = HashMap()
        words.forEach {
            val wordmask = getMask(it, 0)
            wordMaskMap.put(wordmask, wordMaskMap.getOrDefault(wordmask,0)+1)
        }

        var result:MutableList<Int> = mutableListOf()

        puzzles.forEach { puzzle ->
            val puzzleMask = getMask(puzzle, 1)
            var subsetMask = puzzleMask
            var firstCharMask = 1 shl (puzzle[0] - 'a')
            var count = wordMaskMap.getOrDefault(firstCharMask,0)

            while(subsetMask != 0) {
                count += wordMaskMap.getOrDefault(subsetMask or firstCharMask,0)
                subsetMask = (subsetMask - 1) and puzzleMask
            }

            result.add(count)
        }

        return result
    }

    fun getMask(word:String, start:Int): Int {
        var mask = 0
        for (i in start until word.length) {mask = mask or (1 shl (word[i] - 'a'))}
        return mask
    }
}
// @lc code=end

