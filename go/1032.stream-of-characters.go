package leetcode

/*
 * @lc app=leetcode id=1032 lang=golang
 *
 * [1032] Stream of Characters
 */

// @lc code=start
type Trie struct {
	next   map[byte]*Trie
	isWord bool
}

type StreamChecker struct {
	trie *Trie
	q    []byte
	size int
}

func Constructor(words []string) StreamChecker {
	maxSize := 0
	t := &Trie{
		next:   make(map[byte]*Trie),
		isWord: false,
	}

	for _, word := range words {
		insert(word, t)
		if len(word) > maxSize {
			maxSize = len(word)
		}
	}
	return StreamChecker{t, make([]byte, 0), maxSize}
}

func insert(word string, t *Trie) {
	cur := t
	for i := len(word) - 1; i >= 0; i-- {
		curChar := word[i]
		if _,ok := cur.next[curChar]; !ok{
			cur.next[curChar] = &Trie{
				make(map[byte]*Trie),
				false,
			}
		}
		cur = cur.next[curChar]
	}
	cur.isWord = true
}

func (this *StreamChecker) Query(letter byte) bool {
	this.q = append([]byte{letter}, this.q...)
	return this.isWord()
}

func (this *StreamChecker) isWord() bool{
	cur := this.trie
	for i:=0; i<len(this.q); i++{
		curChar := this.q[i]
		if _, ok := cur.next[curChar]; !ok{
			return false
		} else{
			cur = cur.next[curChar]
			if cur.isWord == true{
				return true
			}
		}
	}
	return false
}


/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
// @lc code=end

