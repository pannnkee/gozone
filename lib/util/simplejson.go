package util

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
)

// 结构体转map
func Struct2JsonMap(data interface{}) (map[string]interface{}, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	newJson, err := simplejson.NewJson(bytes)
	if err != nil {
		return nil, err
	}
	jsonMap, err := newJson.Map()
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}
