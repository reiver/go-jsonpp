package jsonpp

import (
	"bytes"
	"io"
)

func FPrettyPrint(writer io.Writer, json []byte) error {
	if nil == writer {
		return errNilWriter
	}

	var buffer bytes.Buffer

	err := jsonPrettyPrint(&buffer, json)
	if nil != err {
		return err
	}

	_, err = buffer.WriteTo(writer)
	if nil != err {
		return err
	}

	return nil
}
