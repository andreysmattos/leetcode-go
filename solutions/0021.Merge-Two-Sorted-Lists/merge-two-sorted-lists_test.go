package leetcode

import (
	"leetcode/structures"
	"reflect"
	"testing"
)

type mergeTestCase struct {
	intList1 []int
	intList2 []int
	expected []int
}

func Test_Problem21(t *testing.T) {

	testCases := []mergeTestCase{
		{
			intList1: []int{1, 2, 4},
			intList2: []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			intList1: []int{},
			intList2: []int{},
			expected: []int{},
		},
		{
			intList1: []int{},
			intList2: []int{0},
			expected: []int{0},
		},
		{
			intList1: []int{1, 3, 5},
			intList2: []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for i, testCase := range testCases {
		result := mergeTwoLists(Ints2List(testCase.intList1), Ints2List(testCase.intList2))
		resultSlice := structures.List2Ints(result)

		if !reflect.DeepEqual(resultSlice, testCase.expected) {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, testCase.expected, resultSlice)
		}
	}
}
