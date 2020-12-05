package json

import (
	"fmt"
	"testing"
)

func Test_GetJsonValue(t *testing.T) {
	str := "{\"remarks\":{\"gBCode\":\"\",\"latitude\":\"\",\"direction\":\"in\",\"longitude\":\"\",\"outputTime\":60000,\"tamperEnable\":0,\"tamperSupport\":0,\"recognitionMode\":\"card,face,fingerprint,qrcode\",\"mediumTypeAbility\":\"card,face,fingerprint,qrcode\"},\"channelNo\":1,\"channelUuid\":\"68418E40DDD4482D94BB02A646849249\",\"multiRecognize\":1,\"mediumTypeAbility\":\"card,face,fingerprint,qrcode\"}"
	value, err := GetJsonValue(str, "remarks.mediumTypeAbility")
	if err != nil {
		fmt.Printf("%v \n", err)
		return
	}
	fmt.Printf("%v \n", value)
}
