package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_http(t *testing.T) {
	arr := []string{"192.168.9.241:31009/2,04696ab7153f", "192.168.9.241:31008/5,046854fc5efe", "192.168.9.241:31009/4,046706a9e0d6", "192.168.9.241:31008/6,0466f21c330d"}

	for _, url := range arr {
		resp, _ := http.Get("http://" + url)
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(len(bytes))
	}

	//resp, _ := http.Get("http://192.168.9.241:33001/v3/file/download?fid=1348866238558126080")
	//resp, _ := http.Get("http://192.168.9.241:31009/2,049c49cbf939")
	//bytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(len(bytes))
}
