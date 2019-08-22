package length_prefixed_string

import (
	"encoding/binary"
	"fmt"
	"io"
)

func readLength(r io.Reader) (uint32, error) {
	var b uint8

	err := binary.Read(r, binary.LittleEndian, &b)
	if err != nil {
		return 0, err
	}

	// 'The high bit MUST be used to indicate that the length continues in the next byte.'
	if b&1 > 128 {
		return 0, fmt.Errorf("not implemented, todo")
	}

	return uint32(b), nil
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/10b218f5-9b2b-4947-b4b7-07725a2c8127
func Read(r io.Reader) (string, error) {
	l, err := readLength(r)
	if err != nil {
		return "", err
	}

	b := make([]byte, l)
	_, err = r.Read(b)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
