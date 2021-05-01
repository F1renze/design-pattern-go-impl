package factory_method

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	f1 := InsertSorterFactory{}
	sorter := f1.CreateSorter()

	tt := []struct {
		In     []int
		Expect []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
	}

	for i := 0; i < len(tt); i++ {
		sorter.Sort(tt[i].In)
		fmt.Println(tt[i].In)
		if !reflect.DeepEqual(tt[i].In, tt[i].Expect) {
			t.Errorf("expect %+v, got %+v", tt[i].Expect, tt[i].In)
		}
	}

}
