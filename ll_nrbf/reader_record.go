package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/954a0657-b901-4813-9398-4ec732fe8b32
const (
	RTE_0_SERIALIZED_STREAM_HEADER     = 0x00
	RTE_5_CLASS_WITH_MEMBERS_AND_TYPES = 0x05
	RTE_6_BINARY_OBJECT_STRING         = 0x06
	RTE_12_BINARY_LIBRARY              = 0x0C
)

type RecordType uint8

type Record struct {
	RecordType RecordType
	Record     interface{}
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/8b313786-0baf-4f01-bc45-3a4c70af3e01#gt_dca3e776-890f-48c8-be62-094a5f2fcf71
func (d *Decoder) NextRecord() (rec Record, err error) {
	// Read the record type
	err = binary.Read(d.r, binary.BigEndian, &rec.RecordType)
	if err != nil {
		return
	}

	fmt.Println("Read record: ", rec.RecordType)

	switch rec.RecordType {
	case RTE_0_SERIALIZED_STREAM_HEADER:
		rec.Record, err = d.decodeSerizlizedStreamHeader()
	case RTE_5_CLASS_WITH_MEMBERS_AND_TYPES:
		rec.Record, err = d.decodeRecordWithMembersAndTypes()
	case RTE_6_BINARY_OBJECT_STRING:
		rec.Record, err = d.decodeBinaryObjectString()
	case RTE_12_BINARY_LIBRARY:
		rec.Record, err = d.decodeRecordBinaryLibrary()
	default:
		err = fmt.Errorf("unknown record type, got %d", rec.RecordType)
	}

	return
}
