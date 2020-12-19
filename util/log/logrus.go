package log

import (
	"github.com/sirupsen/logrus"
	"thinkingingolang/util/log/hook"
)

func AddLineHook() {
	logrus.AddHook(hook.NewLineHook())
}

func AddLogFileHook(path, file string) {
	logrus.AddHook(hook.NewFileRotationHook(path, file))
}
