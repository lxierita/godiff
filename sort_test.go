package sort

import (
	"reflect"
	"testing"
)

func TestSortRecords_should_return_the_arrays_sorted_in_ascending_order(t *testing.T) {
	arr := [5][2]int{{1, 2}, {2, 4}, {3, 6}, {4, 8}, {0, 0}}
	expected := [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

	if sorted := SortRecords(arr); reflect.DeepEqual(sorted, expected) == false {
		t.Errorf("SortRecords(arr) = %v; \n expect %v", sorted, expected)
	}
}
