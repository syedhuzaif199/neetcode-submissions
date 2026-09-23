type MinStack struct {
	data []int
	min []int
	top int
}

func Constructor() MinStack {
	return MinStack {}
}

func (this *MinStack) Push(val int) {
	if this.top < len(this.data) {
		this.data[this.top] = val
	} else {
		this.data = append(this.data, val)
	}
	m := val
	if this.top > 0 {
		m = min(this.min[this.top-1], val)
	}
	if this.top < len(this.min) {
		this.min[this.top] = m
	} else {
		this.min = append(this.min, m)
	}
	this.top += 1
}

func (this *MinStack) Pop() {
	if this.top > 0 {
		this.top -= 1
	}
}

func (this *MinStack) Top() int {
	if this.top > 0 {
		return this.data[this.top-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.top > 0 {
		return this.min[this.top-1]
	}
	return 0
}
