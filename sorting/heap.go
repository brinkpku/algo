package sorting

func heapSort(nums []int) {
	hp := &maxHeap{
		data: nums,
	}
	for i := len(nums) - 1; i >= 0; i-- {
		hp.siftDown(i)
	}
	for i := range nums {
		hp.siftDown2(0, len(nums)-i)
		hp.swap(0, len(nums)-i-1)
	}
}

type maxHeap struct {
	data []int
}

func (mh *maxHeap) left(i int) int {
	return 2*i + 1
}

func (mh *maxHeap) right(i int) int {
	return 2*i + 2
}

func (mh *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (mh *maxHeap) len() int {
	return len(mh.data)
}

func (mh *maxHeap) swap(i, j int) {
	mh.data[i], mh.data[j] = mh.data[j], mh.data[i]
}

func (mh *maxHeap) siftUp(i int) {
	for {
		p := mh.parent(i)
		if p < 0 || mh.data[i] <= mh.data[p] {
			break
		}
		mh.swap(i, p)
		i = p
	}
}

func (mh *maxHeap) siftDown(i int) {
	for {
		l := mh.left(i)
		r := mh.right(i)
		root := i
		if l < mh.len() && mh.data[l] > mh.data[root] {
			root = l
		}
		if r < mh.len() && mh.data[r] > mh.data[root] {
			root = r
		}
		if root == i {
			break
		}
		mh.swap(root, i)
		i = root
	}
}

func (mh *maxHeap) siftDown2(i int, n int) {
	for {
		l, r, root := mh.left(i), mh.right(i), i
		if l < n && mh.data[root] < mh.data[l] {
			root = l
		}
		if r < n && mh.data[root] < mh.data[r] {
			root = r
		}
		if root == i {
			break
		}
		mh.swap(root, i)
		i = root
	}
}

func (mh *maxHeap) push(v int) {
	mh.data = append(mh.data, v)
	mh.siftUp(len(mh.data) - 1)
}

func (mh *maxHeap) pop() int {
	mh.swap(0, len(mh.data)-1)
	v := mh.data[len(mh.data)-1]
	mh.data = mh.data[:len(mh.data)-1]
	mh.siftDown(0)
	return v
}

func (mh *maxHeap) peek() int {
	return mh.data[0]
}
