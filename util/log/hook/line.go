package hook

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

func NewLineHook() *LineHook {
	return new(LineHook)
}

type LineHook struct {
}

func (lh LineHook) Levels() []log.Level {
	return log.AllLevels
}

func (lh LineHook) Fire(entry *log.Entry) error {
	entry.Data["line"] = findCaller(0)
	return nil
}

func findCaller(skip int) string {
	file := ""
	line := 0
	var pc uintptr
	for i := 0; i < 11; i++ {
		file, line, pc = getCaller(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}

	fullFnName := runtime.FuncForPC(pc)
	return fmt.Sprintf("%s:%d:%s()", file, line, fullFnName.Name())
}

func getCaller(skip int) (string, int, uintptr) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0, pc
	}
	n := 0

	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line, pc
}
