import java.util.*;

/*
 * @lc app=leetcode id=1832 lang=java
 *
 * [1832] Check if the Sentence Is Pangram
 */

// @lc code=start
class Solution {
    public boolean checkIfPangram(String sentence) {
        Map<Integer,Integer> map = new HashMap<>();
        sentence.chars().forEach(i -> map.put(i-97, 1));

        return map.size() == 26;
    }
}
// @lc code=end

