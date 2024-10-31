package utils

import "encoding/json"

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func FormatJSON(str string) interface{} {
	var result interface{}
	json.Unmarshal([]byte(str), &result)
	return result
}
