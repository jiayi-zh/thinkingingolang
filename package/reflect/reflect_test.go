package reflect

import (
	"fmt"
	"testing"
)

type User interface {
	Speak(a, b int, c string) (int, int)
}

type City struct {
	CityName string `json:"cityName"`
}

type Person struct {
	Id   int64   `json:"id"`
	Name *string `json:"name"`
	Age  int     `json:"age"`
	*City
}

func (p Person) Speak(a, b int, c string) (int, int) {
	fmt.Println(fmt.Sprintf("a: %d, b: %d, c: %s, detail: %+v", a, b, c, p))
	return 0, 0
}

func Test_reflectApi(t *testing.T) {
	name := "JiaYi"
	p := Person{
		Id:   1,
		Name: &name,
		City: &City{
			CityName: "ShangHai",
		},
	}
	reflectApi(&p)
}
