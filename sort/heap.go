package sort

func adjustHeap(arr SortInterface, index int, length int, isAsc bool) {
	for i := 2*index + 1; i < length; i = 2*i + 1 {
		if i+1 < length {
			if (arr.Compare(i, i+1) < 0 && isAsc) || (arr.Compare(i, i+1) > 0 && !isAsc) {
				i++
			}
		}
		if (arr.Compare(i, index) > 0 && isAsc) || (arr.Compare(i, index) < 0 && !isAsc) {
			arr.Swap(i, index)
			index = i
		} else {
			break
		}
	}
}

// HeapSort is to sort the arr by heap. Not stable sort.
// the parameter - isAsc means the arr should be ordered by ASC or DESC after sorting.
func HeapSort(arr SortInterface, isAsc bool) {

	// create the heap from the first non-leaf node
	for i := arr.Len()/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, arr.Len(), isAsc)
	}

	// swap the first and the last one. then adjust the heap again to find the largest or smallest one.
	for j := arr.Len() - 1; j > 0; j-- {
		arr.Swap(0, j)
		adjustHeap(arr, 0, j, isAsc)
	}
}
