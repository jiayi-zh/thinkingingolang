package hook

import (
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	ByteSize3G = 3 * 1024 * 1024 * 1024
)

func NewFileRotationHook(path, file string) *lfshook.LfsHook {
	writer, err := rotateLogs.New(
		fmt.Sprintf("%s%c%s%s", path, os.PathSeparator, file, ".%Y%m%d%H%M.log"),
		rotateLogs.WithLinkName(fmt.Sprintf("%s%c%s", path, os.PathSeparator, file)),
		rotateLogs.WithRotationSize(ByteSize3G),
		rotateLogs.WithMaxAge(time.Minute*5),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	return lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: false})
}
