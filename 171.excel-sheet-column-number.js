/*
 * @lc app=leetcode id=171 lang=javascript
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
/**
 * @param {string} columnTitle
 * @return {number}
 */
var titleToNumber = function(columnTitle) {
    //split the string into an array of characters
    let arr = columnTitle.split('');
    //create a variable to store the result
    let result = 0;
    //loop through the array
    for (let i = 0; i < arr.length; i++) {
        //convert the character to a number
        let num = arr[i].charCodeAt(0) - 64;
        //multiply the result by 26 and add the number
        result += num * Math.pow(26, arr.length - i - 1);
    }
    return result
};
// @lc code=end

