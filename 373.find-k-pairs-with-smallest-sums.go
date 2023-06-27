import "container/heap"

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start

// Heap is a min heap of []int
type Heap [][]int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(i interface{}) { *h = append(*h, i.([]int)) }
func (h *Heap) Pop() interface{}   { v := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return v }

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &Heap{}
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, []int{nums1[i] + nums2[0], i, 0})
	}
	var res [][]int
	for k > 0 && h.Len() > 0 {
		k--
		v := heap.Pop(h).([]int)
		res = append(res, []int{nums1[v[1]], nums2[v[2]]})
		if v[2] < len(nums2)-1 {
			heap.Push(h, []int{nums1[v[1]] + nums2[v[2]+1], v[1], v[2] + 1})
		}
	}
	return res
}
// @lc code=end

