package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []int{1,4,6,3,5,9,7}
	//BubbleSort(tests, len(tests))
	InsertSort(tests, len(tests))
	fmt.Println(tests)
}
