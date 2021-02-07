package time

import (
	"fmt"
	"testing"
	"time"
)

func Test_TimeNowAddMinute(t *testing.T) {
	unix := time.Now().Unix()

	t2 := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
	fmt.Println(t2)
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
