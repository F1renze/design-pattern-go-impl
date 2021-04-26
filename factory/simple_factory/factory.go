package simple_factory

type Approach int

const (
	BubbleSort Approach = iota
)

type Sorter interface {
	Sort([]int)
}

type BubbleSorter struct {
}

func (b BubbleSorter) Sort(arr []int) {

}

func NewSorter(a Approach) Sorter {
	switch a {
	case BubbleSort:
		return BubbleSorter{}
	}
	return nil
}
