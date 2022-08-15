package leetcode

/*
 * @lc app=leetcode id=1209 lang=golang
 *
 * [1209] Remove All Adjacent Duplicates in String II
 */

// @lc code=start
func removeDuplicates(s string, k int) string {
	stack := MyStack{}

	for _, r := range s {
		stack.Push(r)

		for(stack.TopKElementsEqual(k)) {
			for i := 0; i < k; i++ {
				stack.Pop()
			}
		}
	}

	return string(stack.GetQueue())
}

type MyStack struct {
	queue []rune
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]rune, 0),
	}
}

func (this *MyStack) Push(x rune)  {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() rune {
	ret := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return ret
}

func (this *MyStack) GetQueue() []rune {
	return this.queue
}

func (this *MyStack) TopKElementsEqual(k int) bool {
	if len(this.queue) < k {
		return false
	}
	pos := len(this.queue) - 1
	for i := 0; i < k; i++ {
		if this.queue[pos] != this.queue[pos-i] {
			return false
		}
	}

	return true
}

// @lc code=end

