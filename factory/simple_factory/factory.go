package simple_factory

type Approach int

const (
	BubbleSort Approach = iota
)

type ISorter interface {
	Sort([]int)
}

type BubbleSorter struct {
}

func (b BubbleSorter) Sort(arr []int) {

}

func NewSorter(a Approach) ISorter {
	switch a {
	case BubbleSort:
		return BubbleSorter{}
	}
	return nil
}
