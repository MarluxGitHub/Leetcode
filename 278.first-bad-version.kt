/*
 * @lc app=leetcode id=278 lang=kotlin
 *
 * [278] First Bad Version
 */

// @lc code=start
/* The isBadVersion API is defined in the parent class VersionControl.
      def isBadVersion(version: Int): Boolean = {} */

class Solution: VersionControl() {
    override fun firstBadVersion(n: Int) : Int {
        var low =1
        var high=n

        while (low<high)
        {
            var middel =low +Math.floor((high-low).toDouble()/2).toInt()
            if (!isBadVersion(middel))
                low=middel+1
            else
                high=middel
        }

        return low
    }
}
// @lc code=end

