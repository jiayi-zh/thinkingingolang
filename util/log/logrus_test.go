package log

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_Logrus(t *testing.T) {
	AddLineHook()
	AddLogFileHook("D:\\doc\\temp20201219", "daily")

	go func() {
		for range time.Tick(time.Millisecond * 100) {
			log.WithFields(log.Fields{
				"key1": "v1",
				"key2": "v2",
				"key3": "v3",
			}).Info("hello1")
		}
	}()
	go func() {
		for range time.Tick(time.Millisecond * 100) {
			log.WithFields(log.Fields{
				"key4": "v4",
				"key5": "v5",
				"key6": "v6",
			}).Warnf("hello2")
		}
	}()
	go func() {
		for range time.Tick(time.Millisecond * 100) {
			log.WithFields(log.Fields{
				"key7": "v7",
				"key8": "v8",
				"key9": "v9",
			}).Errorf("hello3")
		}
	}()

	select {}
}
