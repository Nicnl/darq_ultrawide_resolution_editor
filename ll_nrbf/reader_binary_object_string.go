package ll_nrbf

import (
	"encoding/binary"
	"fmt"
	"nrbf/length_prefixed_string"
)

type BinaryObjectString struct {
	ObjectId int32
	Value    string
}

func (d *Decoder) decodeBinaryObjectString() (bos BinaryObjectString, err error) {
	// ObjectId, 'The value MUST be a positive integer'
	err = binary.Read(d.r, binary.LittleEndian, &bos.ObjectId)
	if err != nil {
		return
	}

	if bos.ObjectId < 0 {
		err = fmt.Errorf("invalid ObjectId, expected a positive integer, got %d", bos.ObjectId)
		return
	}

	// Value
	bos.Value, err = length_prefixed_string.Read(d.r)
	if err != nil {
		return
	}

	return
}
