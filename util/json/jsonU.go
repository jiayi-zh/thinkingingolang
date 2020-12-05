package json

import (
	"encoding/json"
	"strings"
)

func GetJsonValue(jsonStr, path string) (interface{}, error) {
	jsonStruct := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonStruct)
	if err != nil {
		return nil, err
	}
	jsonKeyList := strings.Split(path, ".")
	for i, key := range jsonKeyList {
		jsonVal, ok := jsonStruct[key]
		if !ok {
			return nil, nil
		} else {
			if i == len(jsonKeyList)-1 {
				return jsonVal, nil
			} else {
				if v, ok := jsonVal.(map[string]interface{}); ok {
					jsonStruct = v
				} else {
					return nil, nil
				}
			}
		}
	}
	return nil, nil
}
