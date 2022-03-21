/*
 * @lc app=leetcode id=392 lang=javascript
 *
 * [392] Is Subsequence
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function(s, t) {
    let curr = 0;
    for (let i = 0; i < t.length; i++) {
        if (t[i] === s[curr]) curr++;
    }
    if (curr === s.length) return true;
    return false;
};
// @lc code=end

