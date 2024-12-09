package utils

// quickSortStrings sorts a slice of strings lexicographically using the quicksort algorithm.
func QuickSortStrings(data []string) {
	// Base case: if the slice has less than two elements, it is already sorted.
	if len(data) < 2 {
		return
	}
	// Choose a pivot element (in this case, the middle element of the slice).
	pivot := data[len(data)/2]
	left, right := 0, len(data)-1
	// Partition the slice into two subarrays: one with elements less than the pivot and the other with elements greater than the pivot.
	for left <= right {
		// Move the left pointer to the right while the element at the left is less than the pivot.
		for data[left] < pivot {
			left++
		}
		// Move the right pointer to the left while the element at the right is greater than the pivot.
		for data[right] > pivot {
			right--
		}
		// If the left pointer is less than or equal to the right pointer, swap the elements at the left and right.
		if left <= right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}
	// Recursively sort the left and right subarrays.
	QuickSortStrings(data[:right+1]) // Left part.
	QuickSortStrings(data[left:])    // Right part.
}
