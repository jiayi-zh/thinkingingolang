package systemdict

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	fileDir = "D:\\doc\\doc_git\\docs_rd_department\\软件研发相关规范\\广拓软件系统专用名词与字典表.md"
)

func Test_DictImport(t *testing.T) {
	file, err := os.Open(fileDir)
	if err != nil {
		log.Fatalf("open file fail, cause: %v", err)
	}
	defer file.Close()

	var groupId int64
	var groupName, groupType string
	var fd, fh, ff bool
	row := 0
	buf := bufio.NewReader(file)
	for {
		row++
		bytes, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		//log.Info(string(bytes))

		columnList := strings.Split(string(bytes), "|")
		// 8 --> | 类型组ID | 类型组名称 | 类型组字串 | 类型ID | 类型名称 | 类型字串 |
		if len(columnList) != 8 {
			if fd {
				fd = false
			}
			continue
		} else {
			if !fd {
				fd = true
				fh, ff = true, true
			}
		}
		if fh {
			fh = false
			continue
		}
		if ff {
			ff = false
			continue
		}

		if len(strings.TrimSpace(columnList[1])) > 0 {
			groupId, err = strconv.ParseInt(strings.TrimSpace(columnList[1]), 10, 64)
			if err != nil {
				log.Fatalf("row [%d] : data[%s] type error, please check if first", row, string(bytes))
			}
			groupName, groupType = strings.TrimSpace(columnList[2]), strings.TrimSpace(columnList[3])
		}
		var typeId int64
		if len(strings.TrimSpace(columnList[4])) > 0 {
			typeId, err = strconv.ParseInt(strings.TrimSpace(columnList[4]), 10, 64)
			if err != nil {
				log.Fatalf("row [%d] : data[%s] type error, please check if first", row, string(bytes))
			}
		}
		typeName, typeCode := strings.TrimSpace(columnList[5]), strings.TrimSpace(columnList[6])
		log.Infof("row [%d]: 类型组ID: %d, 类型组名称: %s, 类型组字串: %s, 类型ID: %d, 类型名称: %s, 类型字串: %s",
			row, groupId, groupName, groupType, typeId, typeName, typeCode)
	}
}

type DictRow struct {
}
