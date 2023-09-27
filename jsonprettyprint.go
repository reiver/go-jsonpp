package jsonpp

import (
	"bytes"
	"encoding/json"
)

const indent string = "\t"

func jsonPrettyPrint(dst *bytes.Buffer, src []byte) error {
	if nil == dst {
		return errInternalError
	}

	err := json.Indent(dst, src, "", indent)
	if nil != err {
		return err
	}

	return nil
}
