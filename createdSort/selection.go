package createdSort

// SelectionSort 選択ソート
func SelectionSort(s []int) []int {
	for i, _ := range s {
		idx, _ := getMinIndex(s[i:])
		s[i], s[i+idx] = s[i+idx], s[i]
	}

	return s
}

func getMinIndex(s []int) (int, int) {
	n := s[0]
	idx := 0
	for i, v := range s {
		if n > v {
			n = v
			idx = i
		}
	}
	return idx, n
}
