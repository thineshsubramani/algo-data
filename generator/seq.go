package generator

// Seq generates a slice of integers from 1 to size.
func Seq(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i + 1
	}
	return result
}

// SeqReverse generates a slice of integers from size down to 1.
func SeqReverse(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = size - i
	}
	return result
}