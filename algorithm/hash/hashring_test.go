package hash

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_MongoDbTest(t *testing.T) {
	hashRing := NewConsistent()
	for i := 1; i <= 30*12; i++ {
		hashRing.Add(fmt.Sprintf("node%d", i))
	}

	for i := 0; i < 10; i++ {
		node1, _ := hashRing.Get("data1")
		node2, _ := hashRing.Get("data2")
		node3, _ := hashRing.Get("data3")
		node4, _ := hashRing.Get("data4")
		fmt.Println("data1:", node1, ",data2:", node2, ",data3:", node3, ",data4:", node4)
	}

	hashRing.Add(strconv.FormatInt(30*12+1, 10))
	hashRing.Add(strconv.FormatInt(30*12+2, 10))
	hashRing.Add(strconv.FormatInt(30*12+3, 10))
	fmt.Println("-------------------------------------")

	node1, _ := hashRing.Get("data1")
	node2, _ := hashRing.Get("data2")
	node3, _ := hashRing.Get("data3")
	node4, _ := hashRing.Get("data4")
	fmt.Println("data1:", node1, ",data2:", node2, ",data3:", node3, ",data4:", node4)
}
