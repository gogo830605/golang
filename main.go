package main

import (
	// "inwin/routes"
	"fmt"
	// "time"
)

func BuildMaxHeap(in []int) []int {
	heap := []int{}
	for _, value := range in {
		// insert a new node at the bottom
		heap = append(heap, value)
		// do heapify to make the new node to right position
		heap = MaxHeapifyBottomUp(heap)
	}

	return heap
}

func MaxHeapifyBottomUp(in []int) []int {
	// the index of new node, the latest one element in slice
	newELeIndex := len(in) - 1
	for {
		// get the parent node index
		parentIndex := (newELeIndex+1)/2 - 1
		// break if the node is to the top of heap
		if parentIndex < 0 {
			break
		}
		// if parent node < the new node, swap
		if in[newELeIndex] > in[parentIndex] {
			in[parentIndex], in[newELeIndex] = in[newELeIndex], in[parentIndex]
			newELeIndex = parentIndex
		} else {
			// if not, break here
			break
		}
	}
	return in
}

func main() {
	// router.SetupRouter()
	// router := router.SetupRouter()
	// router.Run(":8888")
	// c1 := make(chan string)
    // c2 := make(chan string)

    arr := []int{1, 2, 3, 4, 5, 5, 6, 6, 1, 2, 10, 9, 20}
    fmt.Println(BuildMaxHeap(arr))
}