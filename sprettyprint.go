package jsonpp

import (
	"bytes"
)

func SPrettyPrint(json []byte) string {
	var buffer bytes.Buffer

	err := jsonPrettyPrint(&buffer, json)
	if nil != err {
		return ""
	}

	return buffer.String()
}
