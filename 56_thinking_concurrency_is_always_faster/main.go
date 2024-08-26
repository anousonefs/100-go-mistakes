package main

import (
	"fmt"
)

/* As Go developers, we can’t create threads directly, but we can create goroutines, which can be thought of as application-level threads. However, whereas an OS thread is context-switched on and off a CPU core by the OS, a goroutine is context-switched on and off an OS thread by the Go runtime. Also, compared to an OS thread, a goroutine has a smaller memory footprint: 2 KB for goroutines from Go 1.4. An OS thread depends on the OS, but, for example, on Linux/x86-32, the default size is 2 MB (see http://mng.bz/DgMw). Having a smaller size makes context switching faster. */

/* NOTE Context switching a goroutine versus a thread is about 80% to 90% faster, depending on the architecture. */

/* A goroutine has a simpler lifecycle than an OS thread. It can be doing one of the following: */
/* 1. Executing—The goroutine is scheduled on an M and executing its instructions. */
/* 2. Runnable—The goroutine is waiting to be in an executing state. */
/* 3. Waiting—The goroutine is stopped and pending something completing, such as a system call or a synchronization operation (such as acquiring a mutex). */

// mergeSort is a recursive function that sorts an array using the merge sort algorithm.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Divide the array into two halves.
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the sorted halves.
	return merge(left, right)
}

// merge combines two sorted arrays into a single sorted array.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare elements from both arrays and add the smaller one to the result.
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// If there are remaining elements in the left array, add them to the result.
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// If there are remaining elements in the right array, add them to the result.
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

// slow if we use with small number
/* func parallelMergesortV1(s []int) { */
/* 	if len(s) <= 1 { */
/* 		return */
/* 	} */
/* 	middle := len(s) / 2 */
/* 	var wg sync.WaitGroup */
/* 	wg.Add(2) */
/* 	go func() { */
/* 		defer wg.Done() */
/* 		parallelMergesortV1(s[:middle]) */
/* 	}() */
/* 	go func() { */
/* 		defer wg.Done() */
/* 		parallelMergesortV1(s[middle:]) */
/* 	}() */
/* 	wg.Wait() */
/* 	merge(s, middle) */
/* } */

// fast
/* const max = 2048 */
/* func parallelMergesortV2(s []int) { */
/* 	if len(s) <= 1 { */
/* 		return */
/* 	} */
/* 	if len(s) <= max { */
/* 		sequentialMergesort(s) */
/* 	} else { */
/* 		middle := len(s) / 2 */
/* 		var wg sync.WaitGroup */
/* 		wg.Add(2) */
/* 		go func() { */
/* 			defer wg.Done() */
/* 			parallelMergesortV2(s[:middle]) */
/* 		}() */
/* 		go func() { */
/* 			defer wg.Done() */
/* 			parallelMergesortV2(s[middle:]) */
/* 		}() */
/* 		wg.Wait() */
/* 		merge(s, middle) */
/* 	} */
/* } */

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
