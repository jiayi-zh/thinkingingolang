package context

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func Test_Context(t *testing.T) {
	chn := make(chan int, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			chn <- i
		}
	}()

	for i := range chn {
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Second * 5)
		}(i)
	}

	wg.Wait()
}

// 获取 arr1 与 arr2 对比后, 需要增和删的部分
// arr1 [1,3,5] arr2 [2,3,6] ==>  [2, 6],  [1, 5]
func intersectComplement(arr1, arr2 []int64) ([]int64, []int64) {
	if len(arr1) == 0 {
		return arr2, make([]int64, 0, 0)
	}
	if len(arr2) == 0 {
		return make([]int64, 0, 0), arr1
	}

	maxLen := len(arr1)
	if len(arr2) > len(arr1) {
		maxLen = len(arr2)
	}

	repeatMp := make(map[int64]int8)
	addRes := make([]int64, 0, maxLen)
	delRes := make([]int64, 0, maxLen)

	for _, v1 := range arr1 {
		var flag = false
		for _, v2 := range arr2 {
			if v1 == v2 {
				repeatMp[v1] = 0
				flag = true
				break
			}
		}
		if !flag {
			delRes = append(delRes, v1)
		}
	}
	for _, v2 := range arr2 {
		if _, ok := repeatMp[v2]; !ok {
			addRes = append(addRes, v2)
		}
	}
	return addRes, delRes
}
