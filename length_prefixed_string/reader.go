package length_prefixed_string

import (
	"encoding/binary"
	"io"
)

func Read(r io.Reader) (string, error) {
	var b uint8

	err := binary.Read(r, binary.LittleEndian, &b)
	if err != nil {
		return "", err
	}

	return "", nil
}
