package linkedlist

type Node struct {
	Value int
	Next  *Node
}

func FromSlice(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	head := &Node{
		Value: nums[0],
	}

	current := head

	for _, val := range nums[1:] {
		current.Next = &Node{
			Value: val,
		}

		current = current.Next
	}

	return head
}