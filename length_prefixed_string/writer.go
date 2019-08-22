package length_prefixed_string

import (
	"encoding/binary"
	"fmt"
	"io"
)

func writeLength(length int, w io.Writer) (err error) {
	if length < 0 {
		err = fmt.Errorf("wtf strings smaller than 0 not supposed to happen")
		return
	}
	if length >= 128 {
		err = fmt.Errorf("strings larger than 127 are not implemented, todo")
		return
	}

	err = binary.Write(w, binary.LittleEndian, uint8(length))
	if err != nil {
		return
	}

	return
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/10b218f5-9b2b-4947-b4b7-07725a2c8127
func Write(w io.Writer, str string) (err error) {
	err = writeLength(len(str), w)
	if err != nil {
		return
	}

	_, err = w.Write([]byte(str))
	if err != nil {
		return
	}

	return nil
}
