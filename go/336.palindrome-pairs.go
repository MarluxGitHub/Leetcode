package leetcode

/*
 * @lc app=leetcode id=336 lang=golang
 *
 * [336] Palindrome Pairs
 */

// @lc code=start
type Trie struct {
	palindromesBelow []int
	Children         map[rune]*Trie
	wordEndIndex     int
}

func reverse(w string) string {
	res := ""
	for _, v := range w {
		res = string(v) + res
	}
	return res
}
func (t *Trie) addword(word string, index int) {
	parent := t
	tmpword := reverse(word)
	for i := range tmpword {
		ch := rune(tmpword[i])
		if isPalindrome(word[0 : len(word)-i]) {
			parent.palindromesBelow = append(parent.palindromesBelow, index)
		}
		if child, ok := parent.Children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{Children: make(map[rune]*Trie), wordEndIndex: -1}
			parent.Children[ch] = newChild
			parent = newChild
		}
	}
	parent.wordEndIndex = index
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func Constructor() Trie {
	m := make(map[rune]*Trie)
	return Trie{Children: m, wordEndIndex: -1}
}

func makeTrie(words []string) Trie {
	t := Constructor()
	for k, v := range words {
		t.addword(v, k)
	}
	return t
}
func palindromePairs(words []string) [][]int {
	t := makeTrie(words)
	//fmt.Printf("%v", t)
	var output [][]int
	for i, word := range words {
		candidates := t.getPalindromesForWord(word, i)
		//fmt.Println("word,candidates", word, candidates)
		for len(candidates) > 0 {
			tmp := candidates[0]
			if tmp != i {
				output = append(output, []int{i, tmp})
			}
			candidates = candidates[1:]
		}
	}
	return output
}

func (t *Trie) getPalindromesForWord(word string, index int) []int {
	var output []int
	for len(word) > 0 {
		if t.wordEndIndex >= 0 {
			if isPalindrome(word) {
				output = append(output, t.wordEndIndex)
			}
		}
		if _, ok := t.Children[rune(word[0])]; !ok {
			return output
		}
		t = t.Children[rune(word[0])]
		word = word[1:]
	}

	if t.wordEndIndex >= 0 {
		output = append(output, t.wordEndIndex)
	}
	output = append(output, t.palindromesBelow...)
	return output
}

// @lc code=end

