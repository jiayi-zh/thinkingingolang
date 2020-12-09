package base64

import (
	"fmt"
	"testing"
)

func Test_Base64CodecApi(t *testing.T) {
	base64 := ImagesToBase64("d:\\Pictures\\表情包\\TIM图片20190920180821.gif")
	fmt.Println(len(base64))
}
