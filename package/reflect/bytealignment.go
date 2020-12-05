package reflect

import (
	"fmt"
	"unsafe"
)

// 12
type A struct {
	a int8
	b int32
	c int16
}

// 8
type B struct {
	a int8
	b int16
	c int32
}

//type C struct {
//	a int8
//	b int64
//	c map[string]string
//}

func byteAlignment() {
	a := A{}
	//fmt.Println(unsafe.Offsetof(&a.a),unsafe.Offsetof(&a.b),unsafe.Offsetof(&a.a))
	fmt.Printf("a.a:%d a.b:%d a.c:%d", &a.a, &a.b, &a.c)

	fmt.Println(unsafe.Sizeof(A{}))
	fmt.Println(unsafe.Sizeof(B{}))
}
