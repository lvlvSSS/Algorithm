package sort

// SortInterface is a type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type SortInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Compare returns -1 if the element at the index - i is lesser that the index - j.
	// Compare returns 1 if the element at the index - i is larger than the index - j.
	// Compare returns 0 if the element at the index - i is equal to the index - j.
	Compare(i, j int) int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// IntSlice attaches the methods of Interface to []int
type IntSlice []int

func (data IntSlice) Len() int {
	return len(data)
}
func (data IntSlice) Compare(i, j int) int {
	if data[i] < data[j] {
		return -1
	} else if data[i] > data[j] {
		return 1
	} else {
		return 0
	}
}
func (data IntSlice) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}
