package goroutinepool

import (
	"fmt"
	"testing"
	"time"
)

func Test_GoroutinePool(testing *testing.T) {
	//创建一个Task
	t := NewTask(func() error {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := NewPool(3)

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	//启动协程池p
	p.Run()
}
