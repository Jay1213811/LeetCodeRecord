package main

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	//按照截止日期从小到大排序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	total := 0
	h := &Heap{}
	//循环,
	for _, cours := range courses {
		t := cours[0]
		d := cours[1]
		//判断total天数小于当前课程的deadline,满足,把其加入heap中
		if total+t <= d {
			total += t
			heap.Push(h, t)
			//不能加入的时候,当前课程的最大天数是否小于heap中的最大天数.
			//小于pop heap中的最大值,把当前天加入,更新total
		} else if h.Len() > 0 && h.IntSlice[0] > t {
			total = total - (h.IntSlice[0] - t)
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

//定义大根堆
type Heap struct {
	IntSlice sort.IntSlice
}

func (h Heap) Len() int {
	return len(h.IntSlice)
}
func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}
func (h Heap) Swap(i, j int) {
	h.IntSlice[i], h.IntSlice[j] = h.IntSlice[j], h.IntSlice[i]
}
func (h *Heap) Pop() interface{} {
	old := h.IntSlice
	res := old[len(old)-1]
	h.IntSlice = old[:len(old)-1]
	return res
}

func (h *Heap) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}


