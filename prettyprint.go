package jsonpp

import (
	"os"
)

func PrettyPrint(json []byte) error {
	return FPrettyPrint(os.Stdout, json)
}
