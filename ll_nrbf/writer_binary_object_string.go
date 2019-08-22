package ll_nrbf

import (
	"encoding/binary"
	"fmt"
	"nrbf/length_prefixed_string"
)

func (e *Encoder) encodeBinaryObjectString(bos BinaryObjectString) (err error) {
	// ObjectId, 'The value MUST be a positive integer'
	if bos.ObjectId < 0 {
		err = fmt.Errorf("invalid ObjectId, expected a positive integer, got %d", bos.ObjectId)
		return
	}

	err = binary.Write(e.w, binary.LittleEndian, bos.ObjectId)
	if err != nil {
		return
	}

	// Value
	err = length_prefixed_string.Write(e.w, bos.Value)
	if err != nil {
		return
	}

	return
}
