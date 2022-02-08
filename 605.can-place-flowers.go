package leetcode

/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
    m := 0

	for pos, flower := range flowerbed {
		if flower == 0 {
			if (pos == 0 || flowerbed[pos-1] == 0) && (pos == len(flowerbed)-1 || flowerbed[pos+1] == 0) {
				flowerbed[pos] = 1
				m++
			}
		}

		if(m >= n){
			return true
		}
	}

	return false
}
// @lc code=end

