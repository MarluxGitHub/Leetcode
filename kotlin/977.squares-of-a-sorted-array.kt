/*
 * @lc app=leetcode id=977 lang=kotlin
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
class Solution {
    fun sortedSquares(A: IntArray): IntArray {
        // Create markers to use to navigate inward since we know that
        // the polar ends are (possibly, but not always) the largest
        var leftMarker = 0
        var rightMarker = A.size - 1

        // Create a marker to track insertions into the new array
        var resultIndex = A.size - 1
        val result = IntArray(A.size)

        // Iterate over the items until the markers reach each other.
        // Its likely a little faster to consider the case where the left
        // marker is no longer producing elements that are less than zero.
        while (leftMarker <= rightMarker) {
            // Grab the absolute values of the elements at the respective
            // markers so they can be compared and inserted into the right
            // index.
            val left = Math.abs(A[leftMarker])
            val right = Math.abs(A[rightMarker])

            // Do checks to decide which item to insert next.
            result[resultIndex] = if (right > left) {
                rightMarker--
                right * right
            } else {
                leftMarker++
                left * left
            }

            // Once the item is inserted we can update the index we want
            // to insert at next.
            resultIndex--
        }

        return result
    }
}
// @lc code=end

