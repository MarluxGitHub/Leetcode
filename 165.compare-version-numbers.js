/*
 * @lc app=leetcode id=165 lang=javascript
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
/**
 * @param {string} version1
 * @param {string} version2
 * @return {number}
 */
var compareVersion = function(version1, version2) {
    var version1Array = version1.split('.');
    var version2Array = version2.split('.');

    let c = 1;

    if (version1Array.length < version2Array.length) {
        version2Array = version1.split('.');
        version1Array = version2.split('.');
        c = -1;
    }

    for (let index = 0; index < version1Array.length; index++) {
        let revision = parseInt((version1Array[index]));

        let revision2 = 0;
        if (version2Array[index] !== undefined) {
            revision2 = parseInt(version2Array[index]);
        }

        if(revision === revision2) continue;

        return (revision > revision2 ? 1 : -1) * c;
    }

    return 0;
};

// @lc code=end

