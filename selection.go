package selection

import "sort"

// Interface is alias of sort.Interface
type Interface = sort.Interface

// Select performs introselect algorithm on data and return index of the kth-smallest value.
func Select(data Interface, k int) int {
	if data.Len() > 0 {
		return introselect(data, 0, data.Len()-1, k)
	}
	return -1
}

func introselect(data Interface, left, right, k int) int {
	if left == right {
		return left
	}
	pivotIndex := pivot(data, left, right)
	pivotIndex = partition(data, left, right, pivotIndex)
	if k == pivotIndex {
		return k
	} else if k < pivotIndex {
		return introselect(data, left, pivotIndex-1, k)
	} else {
		return introselect(data, pivotIndex+1, right, k)
	}
}

func pivot(data Interface, left, right int) int {
	return (left + right) >> 1
}

func partition(data Interface, left, right, pivotIndex int) int {
	data.Swap(pivotIndex, right)
	storeIndex := left
	for i := left; i < right; i++ {
		if data.Less(i, pivotIndex) {
			data.Swap(storeIndex, i)
			storeIndex++
		}
	}
	data.Swap(right, storeIndex)
	return storeIndex
}
