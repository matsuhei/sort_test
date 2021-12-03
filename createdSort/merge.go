package createdSort

// MergeSort マージソート
func MergeSort(s []int) []int {
	var result []int
	// 再起処理の都合上サイズが1以下はソートなしのため即return
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2
	r := MergeSort(s[:mid])
	l := MergeSort(s[mid:])
	i, j := 0, 0

	// ソート部分
	for i < len(r) && j < len(l) {
		if r[i] > l[j] {
			result = append(result, l[j])
			j++
		} else {
			result = append(result, r[i])
			i++
		}
	}

	// 配列マージ部分
	result = append(result, r[i:]...)
	result = append(result, l[j:]...)

	return result
}
