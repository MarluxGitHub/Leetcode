package leetcode

import "sort"

/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */

// @lc code=start
func topKFrequent(words []string, k int) []string {
    	wordFreq := map[string]int{}
	for _, word := range words {
		wordFreq[word]++
	}

	freqWord := map[int][]string{}
	for word, freq := range wordFreq {
		freqWord[freq] = append(freqWord[freq], word)
	}

	res := []string{}
	for i := len(words); i >= 0 && len(res) < k; i-- {
		words := freqWord[i]
		if len(words) == 0 {
			continue
		}
		sort.Strings(words)
		res = append(res, words...)
	}

	return res[:k]
}
// @lc code=end

