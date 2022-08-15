package main

/*
 * @lc app=leetcode id=1461 lang=golang
 *
 * [1461] Check If a String Contains All Binary Codes of Size K
 */

// @lc code=start

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	foundBinaryCodes := make(map[string]struct{}, 0)

	for i := 0; i <= len(s)-k; i++ {
		binaryCode := s[i : i+k]
		foundBinaryCodes[binaryCode] = struct{}{}

		if len(foundBinaryCodes) == 1<<k {
			return true
		}
	}

	return false
}

// @lc code=end

