package stack

type SortedStack struct {
	stack    []int
	auxStack []int
}

func Constructor() SortedStack {
	return SortedStack{
		stack:    []int{},
		auxStack: []int{},
	}
}

func (ss *SortedStack) Push(val int) {
	for len(ss.stack) > 0 && ss.stack[len(ss.stack)-1] < val {
		ss.auxStack = append(ss.auxStack, ss.stack[len(ss.stack)-1])
		ss.stack = ss.stack[:len(ss.stack)-1]
	}
	ss.stack = append(ss.stack, val)
	for len(ss.auxStack) > 0 {
		ss.stack = append(ss.stack, ss.auxStack[len(ss.auxStack)-1])
		ss.auxStack = ss.auxStack[:len(ss.auxStack)-1]
	}
}

func (ss *SortedStack) Pop() {
	if len(ss.stack) > 0 {
		ss.stack = ss.stack[:len(ss.stack)-1]
	}
}

func (ss *SortedStack) Peek() int {
	if len(ss.stack) == 0 {
		return -1
	}
	return ss.stack[len(ss.stack)-1]
}

func (ss *SortedStack) IsEmpty() bool {
	return len(ss.stack) == 0
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
