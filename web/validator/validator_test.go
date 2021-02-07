package validator

import (
	"fmt"
	"testing"
)

type User struct {
	Remarks *string `json:"remarks" validate:"omitempty,lte=2"`
}

func Test_validatorApi(t *testing.T) {
	//str := "rr"
	//user := &User{
	//	Remarks: &str,
	//}
	//
	//vd := validator.New()
	//err := vd.Struct(user)
	//if err != nil {
	//	if _, ok := err.(*validator.InvalidValidationError); ok {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	for _, err := range err.(validator.ValidationErrors) {
	//		fmt.Println(err.Namespace())
	//		fmt.Println(err.Field())
	//		fmt.Println(err.StructNamespace())
	//		fmt.Println(err.StructField())
	//		fmt.Println(err.Tag())
	//		fmt.Println(err.ActualTag())
	//		fmt.Println(err.Kind())
	//		fmt.Println(err.Type())
	//		fmt.Println(err.Value())
	//		fmt.Println(err.Param())
	//		fmt.Println()
	//	}
	//}

	var arr []int
	fmt.Println(len(arr))
}
