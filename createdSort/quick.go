package createdSort

// QuickSort クイックソート
func QuickSort(s []int) []int {
	// 再起処理の都合上サイズが1以下はソートなしのため即return
	if len(s) < 2 {
		return s
	}

	pivot := s[0]
	place := 0

	for j := 0; j < len(s)-1; j++ {
		if s[j+1] < pivot {
			s[j+1], s[place+1] = s[place+1], s[j+1]
			place++
		}
	}
	s[0], s[place] = s[place], s[0]

	first := QuickSort(s[:place])
	second := QuickSort(s[place+1:])
	first = append(first, s[place])

	first = append(first, second...)
	return first
}
