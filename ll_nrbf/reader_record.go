package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

const (
	RECORD_BINARY_LIBRARY = 0x0C
)

type Record struct {
	RecordType uint8
	Data       interface{}
}

func (d *Decoder) NextRecord() (rec Record, err error) {
	// Read the record type
	err = binary.Read(d.r, binary.BigEndian, &rec.RecordType)
	if err != nil {
		return
	}

	switch rec.RecordType {
	case RECORD_BINARY_LIBRARY:
		rec.Data, err = d.decodeRecordBinaryLibrary()
	default:
		err = fmt.Errorf("unknown record type, got %d", rec.RecordType)
	}

	return
}
