package handler

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueueInit(t *testing.T) {
	items := map[string]int64{
		"objectid_1": 1592176571433757000,
		"objectid_2": 1592176571433757001,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, timestamp := range items {
		pq[i] = &Item{
			Value:     value,
			Timestamp: timestamp,
			Index:     1,
		}
		i++
	}
	heap.Init(&pq)
}

func TestPriorityQueuePush(t *testing.T) {
	pq := make(PriorityQueue, 0)
	item := &Item{
		Value:     "objectid_1",
		Timestamp: 1592176571433757000,
	}
	heap.Init(&pq)
	heap.Push(&pq, item)
}

func TestPriorityQueuePop(t *testing.T) {
	items := map[string]int64{
		"objectid_1": 1592176571433757000,
		"objectid_3": 1,
		"objectid_2": 1592176571433757001,
		"objectid_4": 3,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, timestamp := range items {
		pq[i] = &Item{
			Value:     value,
			Timestamp: timestamp,
			Index:     1,
		}
		i++
	}
	heap.Init(&pq)

	a := pq[0]

	fmt.Println("a", a.Value)

	item := heap.Pop(&pq).(*Item)
	if item.Timestamp != 1 {
		t.Errorf("Pop(%q) != %q, want %q", item.Timestamp, 1, 1)
	}
}

func TestPriorityQueuePeak(t *testing.T) {
	items := map[string]int64{
		"objectid_1": 1592176571433757000,
		"objectid_3": 1,
		"objectid_2": 1592176571433757001,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, timestamp := range items {
		pq[i] = &Item{
			Value:     value,
			Timestamp: timestamp,
			Index:     1,
		}
		i++
	}
	heap.Init(&pq)

	item := *pq[0]

	if item.Timestamp != 1 {
		t.Errorf("Peak(%q) != %q, want %q", item.Timestamp, 1, 1)
	}
}
