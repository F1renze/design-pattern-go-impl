package factory_method

type ISorter interface {
	Sort([]int)
}

type InsertSorter struct {
}

func (s InsertSorter) Sort(arr []int) {
	var i, j, tmp int
	for i = 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			tmp = arr[i]
			for j = i-1; j>=0&&arr[j]>tmp; j-=1 {
				arr[j+1] = arr[j]
			}
			j+=1
			arr[j] = tmp
		}
	}
}

type SelectSorter struct {
}

func (s SelectSorter) Sort(arr []int) {
	var i, j, min int
	n := len(arr)
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

// 工厂类
type IFactory interface {
	CreateSorter() ISorter
}

type InsertSorterFactory struct {
}

func (f InsertSorterFactory) CreateSorter() ISorter {
	return InsertSorter{}
}

type SelectSorterFactory struct {
}

func (s SelectSorterFactory) CreateSorter() ISorter {
	return SelectSorter{}
}
