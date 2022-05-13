package leetcode

/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	queue []int
}


func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}


func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}


func (this *MyStack) Pop() int {
	ret := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return ret
}


func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}


func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

