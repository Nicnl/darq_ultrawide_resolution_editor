package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

const (
	RECORD_BINARY_LIBRARY               = 0x0C
	RECORD_CLASS_WITH_MEMBERS_AND_TYPES = 0x05
)

type Record struct {
	RecordType uint8
	Record     interface{}
}

func (d *Decoder) NextRecord() (rec Record, err error) {
	// Read the record type
	err = binary.Read(d.r, binary.BigEndian, &rec.RecordType)
	if err != nil {
		return
	}

	switch rec.RecordType {
	case RECORD_BINARY_LIBRARY:
		rec.Record, err = d.decodeRecordBinaryLibrary()
	case RECORD_CLASS_WITH_MEMBERS_AND_TYPES:
		rec.Record, err = d.decodeRecordWithMembersAndTypes()
	default:
		err = fmt.Errorf("unknown record type, got %d", rec.RecordType)
	}

	return
}
