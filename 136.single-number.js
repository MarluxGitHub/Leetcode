/*
 * @lc app=leetcode id=136 lang=javascript
 *
 * [136] Single Number
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
 var singleNumber = function(nums) {
    var result = nums[0];
    for(var i = 1; i < nums.length; i++) {
        result = result^nums[i];
    }
    return result;
};

// @lc code=end

