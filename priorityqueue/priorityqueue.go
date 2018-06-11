package priorityqueue

import "fmt"

type PriorityQueue interface {
	CompareTo(x PriorityQueue) bool
}

type Queue struct {
	queue []PriorityQueue
}

func NewPriorityQueue() *Queue {
	return &Queue{}
}

func (p *Queue) Push(x PriorityQueue) {
	p.queue = append(p.queue, x)
	idx := p.Length() - 1
	parentIdx := getParent(idx)
	for idx != parentIdx && p.queue[idx].CompareTo(p.queue[parentIdx]) {
		p.swap(idx, parentIdx)
		idx = parentIdx
		parentIdx = getParent(parentIdx)
	}
}

func (p *Queue) Pop() (PriorityQueue, error) {
	if p.Length() == 0 {
		return nil, fmt.Errorf("index out of bound exception")
	}
	p.swap(0, p.Length()-1)
	toPop := p.queue[p.Length()-1]
	p.queue = p.queue[:p.Length()-1]
	p.maxHeapify(0)
	return toPop, nil
}

func (p *Queue) Peek() (PriorityQueue, error) {
	if p.Length() == 0 {
		return nil, fmt.Errorf("index out of bound exception")
	}
	return p.queue[0], nil
}

func (p *Queue) maxHeapify(idx int) {
	left := leftChild(idx)
	right := rightChild(idx)
	min := idx
	if left <= p.Length()-1 && p.queue[left].CompareTo(p.queue[idx]) {
		min = left
	}
	if right <= p.Length()-1 && p.queue[right].CompareTo(p.queue[min]) {
		min = right
	}
	if min != idx {
		p.swap(min, idx)
		p.maxHeapify(min)
	}
}

func (p *Queue) Length() int {
	return len(p.queue)
}

func getParent(i int) int {
	if i%2 == 1 {
		return i / 2
	}
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (p *Queue) swap(i, j int) {
	p.queue[i], p.queue[j] = p.queue[j], p.queue[i]
}
