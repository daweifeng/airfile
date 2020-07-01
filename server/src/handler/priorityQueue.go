package handler

import "container/heap"

type Item struct {
	Value     string
	Timestamp int64
	Index     int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Timestamp < pq[j].Timestamp
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = j
	pq[j].Index = i
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	prev := *pq
	n := len(prev)
	item := prev[n-1]
	prev[n-1] = nil
	*pq = prev[0 : n-1]
	item.Index = -1
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, timestamp int64) {
	item.Value = value
	item.Timestamp = timestamp
	heap.Fix(pq, item.Index)
}
