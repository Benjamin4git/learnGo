package mergesort

type MergeSort struct {
	Value []int
}

func (m *MergeSort) Sort() ([]int, error) {
	m.Value = mergeSort(m.Value)
	return m.Value, nil
}

func mergeSort(r []int) []int {
	len := len(r)
	if len <= 1 {
		return r
	}
	mid := len / 2
	left := mergeSort(r[:mid])
	right := mergeSort(r[mid:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
