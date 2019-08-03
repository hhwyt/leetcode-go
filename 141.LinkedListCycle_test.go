package problem141

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func List(nums []int) *ListNode {
	dummy := &ListNode{}
	node := dummy

	for _, v := range nums {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}

	return dummy.Next
}

func CycleList(nums []int, pos int) *ListNode {
	list := List(nums)
	if pos == -1 {
		return list
	}

	node := list
	for pos > 0 {
		node = node.Next
		pos--
	}

	tail := node
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	return list
}

var testDataSet = []struct {
	nums    []int
	pos     int
	isCycle bool
}{
	{
		[]int{3, 2, 0, -4},
		1,
		true,
	},
	{
		[]int{1, 2},
		1,
		true,
	},
	{
		[]int{1},
		-1,
		false,
	},
	{
		[]int{},
		-1,
		false,
	},
}

func Test_hasCycle(t *testing.T) {
	ast := assert.New(t)
	for _, td := range testDataSet {
		head := CycleList(td.nums, td.pos)
		ast.Equal(td.isCycle, hasCycle(head))
	}
}

func Benchmark_hasCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, td := range testDataSet {
			head := CycleList(td.nums, td.pos)
			hasCycle(head)
		}
	}
}
