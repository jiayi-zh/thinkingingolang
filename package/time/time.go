package time

import (
	"fmt"
	"time"
)

func TimeNowAddMinute(durMin time.Duration) string {
	return time.Now().Add(time.Minute * durMin).Format("2006-01-02 15:04:05")
}

func TimeSchedule() {
	go func() {
		for {
			now := time.Now()
			next := now.Add(time.Second * 3)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

			timer := time.NewTimer(next.Sub(now))
			<-timer.C
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}()
}
