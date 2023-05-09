/*
 * @lc app=leetcode id=649 lang=golang
 *
 * [649] Dota2 Senate
 */

// @lc code=start
func predictPartyVictory(senate string) string {

	i := 0

	for true {
		if checkIfStringOnlyContainsSameChar(senate) {
			break
		}

		c := ""

		if senate[i] == 'R' {
			c = "D"
		} else {
			c = "R"
		}

		senate, ok := deleteFirstCharacterCInStringAfterIndexI(senate, c, i)

		if !ok {
			deleteFirstCharacterCInStringAfterIndexI(senate, c, 0)
		}

		i++

		if i == len(senate) {
			i = 0
		}
	}

	if senate[0] == 'R' {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func checkIfStringOnlyContainsSameChar(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func deleteFirstCharacterCInStringAfterIndexI(s string ,c rune,i int ) (string,bool) {
	for j := i; j < len(s); j++ {
		if string(s[j]) == c {
			return s[:j] + s[j+1:], true
		}
	}
	return s, false
}


// @lc code=end

