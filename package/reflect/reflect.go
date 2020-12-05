package reflect

import (
	"fmt"
	"reflect"
)

func reflectApi(param interface{}) {
	if param == nil {
		return
	}

	// 类型
	rt := reflect.TypeOf(param)

	ptrFlag := rt.Kind() == reflect.Ptr
	fmt.Println("参数类型:", rt.Kind(), "是否为指针类型:", ptrFlag)
	if ptrFlag {
		// Elem(): 返回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice, 会panic
		rt = rt.Elem()
	}
	fmt.Println("类型:", rt.Kind(), "类型名称:", rt.Name(), "类型包路径:", rt.PkgPath(), "类型的字符串表示:", rt.String())
	fmt.Println("存储该类型对象需要字节数:", rt.Size(), "从内存中申请一个该类型值时, 会对齐的字节数:", rt.Align(), "该类型作为结构体字段时, 对齐的字节数:", rt.FieldAlign())
	fmt.Println("结构体字段数:", rt.NumField(), "字段详情如下:")
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		t := field.Type
		if field.Type.Kind() == reflect.Ptr {
			t = field.Type.Elem()
		}
		fmt.Println("字段:", field.Name, "type:", t.Name(), "pkgPath:", t.PkgPath(), "tag:", field.Tag,
			"offset:", field.Offset, "是否为内嵌字段:", field.Anonymous, "index sequence:", field.Index)
	}
	fmt.Println("结构体方法数:", rt.NumMethod(), "方法详情如下:")
	for i := 0; i < rt.NumMethod(); i++ {
		method := rt.Method(i)
		fmt.Println("方法:", method.Name, "type:", method.Type.Name(), "pkgPath:", method.PkgPath, "index:", method.Index)
	}

	// 值
	rv := reflect.ValueOf(param)
	returns := rv.Method(0).Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf("str")})
	fmt.Printf("%v", returns)
}
