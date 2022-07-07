//go:build js && wasm

package filer

import (
	"syscall/js"
)

var (
	bytesArray = js.Global().Get("Uint8Array")
)

func Bytes(bytes []byte) js.Value {
	buf := bytesArray.New(len(bytes))
	for i, b := range bytes {
		buf.SetIndex(i, b)
	}
	return buf
}

func NewBuffer(buffer []byte) js.Value {
	return Buffer.Call("from", Bytes(buffer))
}
