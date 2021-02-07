package http

import (
	"bytes"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
)

func Test_MultipartFormData(t *testing.T) {
	client := resty.New()
	client.SetHostURL("192.168.9.27:9333")

	fileBytes, err := ioutil.ReadFile("d:\\Pictures\\表情包\\manifest.json")
	log.Infof("%v %d", err, len(fileBytes))

	//fileJson := []byte("{\"name\":\"sy.jpg\",\"mime\":\"image/jpeg\",\"size\":93276,\"chunks\":[{\"fid\":\"4,0bc1f2d1a8\",\"offset\":0,\"size\":85220},{\"fid\":\"3,0c3d1ced8e\",\"offset\":85220,\"size\":8056}]}")
	mergeFile := &MergeFile{
		Name: "sy.jpg",
		Mime: "image/jpeg",
		Size: 93276,
		Chunks: []MergeChunkInfo{
			{
				Fid:    "4,0bc1f2d1a8",
				Offset: 0,
				Size:   85220,
			},
			{
				Fid:    "3,0c3d1ced8e",
				Offset: 85220,
				Size:   8056,
			},
		},
	}
	byteJson, err := json.Marshal(mergeFile)
	log.Infof("%v %v", len(byteJson), err)

	buf := new(bytes.Buffer)
	read, err := buf.Write(byteJson)
	log.Infof("%v %v", read, err)

	rsp, err := client.R().
		SetFileReader("file", "manifest.json", buf).
		SetHeader("Content-Type", "application/json").
		Post("http://192.168.9.27:9080/4,14278428cd?cm=true")
	log.Infof("%v %v", err, rsp)
}

type MergeFile struct {
	Name   string           `json:"name"`
	Mime   string           `json:"mime"`
	Size   int64            `json:"size"`
	Chunks []MergeChunkInfo `json:"chunks"`
}

type MergeChunkInfo struct {
	Fid    string `json:"fid"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
}
