package createdSort

// InsertionSort 挿入ソート
func InsertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		j := i - 1
		tmp := s[i]
		for j > -1 && s[j] > tmp {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = tmp
	}

	return s
}
