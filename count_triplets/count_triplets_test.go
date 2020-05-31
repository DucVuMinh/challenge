package count_triplets

import (
	"fmt"
	"testing"
)

type Result struct {
	TotalCount   int64
	FullfilCount int64
	PrevCount    int64
}

func countTriplets(arr []int64, r int64) int64 {
	length := len(arr)
	var totalCount int64
	if length < 3 {
		return 0
	}
	singleMap := map[int64]*Result{}
	for i := 0; i < length; i++ {
		if i == 0 {
			res := Result{1, 0, 0}
			singleMap[arr[i]*r] = &res
			continue
		}
		mulRes := arr[i] * r
		res, ok := singleMap[mulRes]
		if !ok {
			res = &Result{0, 0, 0}

		}
		// count all prev fullfil, with the latest Triplets element index = i
		prev, ok := singleMap[arr[i]]
		if ok {
			totalCount += prev.PrevCount
			res.FullfilCount += prev.PrevCount
			res.PrevCount += prev.TotalCount
		}
		res.TotalCount += 1
		singleMap[mulRes] = res
	}
	return totalCount
}

func TestCount(t *testing.T) {
	c := countTriplets([]int64{1,5,5,25,25,5,25,125}, 5)
	fmt.Println(c)
	c2 := countTriplets([]int64{1,100,100,100}, 1)
	fmt.Println(c2)
}
