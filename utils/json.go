package utils

import "encoding/json"

func JsonMarshal(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}