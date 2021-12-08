package main

import "sort"

func findContentChildren(children []int, biscuits []int) int {
	sort.Ints(children)
	sort.Ints(biscuits)
	get_biscuits_children:=0
	for i,j:=0,0;i<len(children) && j<len(biscuits);{
		if biscuits[j]>=children[i]{
			get_biscuits_children++
			i++
			j++
		}else{
			j++
		}
	}
	return get_biscuits_children
}
