package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sort/createdSort"
	"strconv"
	"time"
)

type results struct {
	start  int64
	end    int64
	result int64
	name   string
}

type sortMethod struct {
	raw []int
}

func main() {
	records := [][]string{
		{
			"count",
			"bubbleResult",
			"selectionResult",
			"insertionSortResult",
			"mergeSortResult",
			"quickSortResult",
			"goSort",
		},
	}

	for i := 0; i < 1000; i += 100 {
		num := 200 * i

		var bResult int64 = 0
		for j := 0; j < 30; j++ {
			s := genSlice(1, 1000000, num, int64(j))
			var bubbleSlice = make([]int, len(s))
			copy(bubbleSlice, s)
			bubbleResult := verify(createdSort.BubbleSort, bubbleSlice, "BubbleSort")
			bResult += bubbleResult.result
			fmt.Printf("bubbleResult time: %d[ns]\n", bubbleResult.result)
		}
		bResult = bResult / 30

		var sResult int64 = 0
		for j := 0; j < 30; j++ {
			s := genSlice(1, 1000000, num, int64(j))
			var selectionSlice = make([]int, len(s))
			copy(selectionSlice, s)
			selectionResult := verify(createdSort.SelectionSort, selectionSlice, "SelectionSort")
			sResult += selectionResult.result
			fmt.Printf("selectionResult time: %d[ns]\n", selectionResult.result)
		}
		sResult = sResult / 30

		var iResult int64 = 0
		for j := 0; j < 30; j++ {
			s := genSlice(1, 1000000, num, int64(j))
			var insertionSortSlice = make([]int, len(s))
			copy(insertionSortSlice, s)
			insertionSortResult := verify(createdSort.InsertionSort, insertionSortSlice, "InsertionSort")
			fmt.Printf("insertionSortResult time: %d[ns]\n", insertionSortResult.result)
			iResult += insertionSortResult.result
		}
		iResult = iResult / 30

		var mResult int64 = 0
		for j := 0; j < 30; j++ {
			s := genSlice(1, 1000000, num, int64(j))
			var mergeSortSlice = make([]int, len(s))
			copy(mergeSortSlice, s)
			mergeSortResult := verify(createdSort.MergeSort, mergeSortSlice, "MergeSort")
			fmt.Printf("mergeSortResult time: %d[ns]\n", mergeSortResult.result)
			mResult += mergeSortResult.result
		}
		mResult = mResult / 30

		var qResult int64 = 0
		for j := 0; j < 30; j++ {
			s := genSlice(1, 1000000, num, int64(j))
			var quickSortSlice = make([]int, len(s))
			copy(quickSortSlice, s)
			quickSortResult := verify(createdSort.QuickSort, quickSortSlice, "QuickSort")
			fmt.Printf("quickSortResult time: %d[ns]\n", quickSortResult.result)
			qResult += quickSortResult.result
		}
		qResult = qResult / 30

		var gResult int64 = 0
		for j := 0; j < 30; j++ {
			s := genSlice(1, 1000000, num, int64(j))
			var goSort = make([]int, len(s))
			copy(goSort, s)
			start := time.Now().UnixNano()
			sort.Sort(sortMethod{raw: goSort})
			end := time.Now().UnixNano()
			goSortResult := results{name: "goSort", start: start, end: end, result: end - start}
			fmt.Printf("goSortResult time: %d[ns]\n", goSortResult.result)
			gResult += goSortResult.result
		}
		gResult = gResult / 30

		records = append(records, []string{
			strconv.Itoa(num),
			strconv.FormatInt(bResult, 10),
			strconv.FormatInt(sResult, 10),
			strconv.FormatInt(iResult, 10),
			strconv.FormatInt(mResult, 10),
			strconv.FormatInt(qResult, 10),
			strconv.FormatInt(gResult, 10),
		})
	}

	fl, _ := os.Create("result.csv")
	defer fl.Close()
	// 書込み
	wr := csv.NewWriter(fl)

	// デリミタ(TSVなら\t, CSVなら,)設定
	wr.Comma = ','

	//デフォルトはLFのみ
	wr.UseCRLF = true

	// ファイルに書き出す
	for _, record := range records {
		if err := wr.Write(record); err != nil {
			// 書き込みエラー発生
			fmt.Println("書き込みエラー発生: ", err)
			return
		}
		// このタイミングで書込みの実施
		wr.Flush()
	}
}

func verify(sortFunc func([]int) []int, s []int, name string) results {
	result := results{}
	result.start = time.Now().UnixNano()
	sortFunc(s)
	result.end = time.Now().UnixNano()
	result.name = name
	result.result = result.end - result.start
	return result
}

// 乱数群生成
func genSlice(min int, max int, num int, seedNum int64) []int {
	numRange := max - min

	exists := make(map[int]bool)
	rand.Seed(seedNum)
	var keys = make([]int, 0, num)
	for count := 0; count < num; {
		n := rand.Intn(numRange) + min
		if exists[n] == false {
			exists[n] = true
			keys = append(keys, n)
			count++
		}
	}

	return keys
}

func (s sortMethod) Len() int {
	return len(s.raw)
}

func (s sortMethod) Less(i, j int) bool {
	return s.raw[i] < s.raw[j]
}

func (s sortMethod) Swap(i, j int) {
	s.raw[i], s.raw[j] = s.raw[j], s.raw[i]
	return
}
