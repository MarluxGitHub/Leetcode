/*
 * @lc app=leetcode id=804 lang=kotlin
 *
 * [804] Unique Morse Code Words
 */

// @lc code=start
class Solution {
    fun uniqueMorseRepresentations(words: Array<String>): Int {
        var map = hashMapOf<Char, String>(
            Pair('a', ".-"),
            Pair('b', "-..."),
            Pair('c', "-.-."),
            Pair('d', "-.."),
            Pair('e', "."),
            Pair('f', "..-."),
            Pair('g', "--."),
            Pair('h', "...."),
            Pair('i', ".."),
            Pair('j', ".---"),
            Pair('k', "-.-"),
            Pair('l', ".-.."),
            Pair('m', "--"),
            Pair('n', "-."),
            Pair('o', "---"),
            Pair('p', ".--."),
            Pair('q', "--.-"),
            Pair('r', ".-."),
            Pair('s', "..."),
            Pair('t', "-"),
            Pair('u', "..-"),
            Pair('v', "...-"),
            Pair('w', ".--"),
            Pair('x', "-..-"),
            Pair('y', "-.--"),
            Pair('z', "--..")
        )

        var morseSet = mutableSetOf<String>()

        for (word in words) {
            var morse = ""
            word.forEachIndexed { _, char ->
                morse += map[char]
             }
             morseSet.add(morse)
        }

        return morseSet.size
    }
}
// @lc code=end

