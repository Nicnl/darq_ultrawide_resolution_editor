package nrbf

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Decoder struct {
	r   io.Reader
	buf []byte
}

type Header struct {
	RecordTypeEnumeration uint8
	ObjectId              int32
	HeaderId              int32
	MajorVersion          int32
	MinorVersion          int32
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/a7e578d3-400a-4249-9424-7529d10d1b3c
// [00] [01 00 00 00] [FF FF FF FF] [01 00 00 00] [00 00 00 00]
// Respectively: RecordTypeEnumeration, ObjectId, HeaderId, MajorVersion and MinorVersion
func (d *Decoder) ParseHeader() (Header, error) {
	var h Header

	// RecordTypeEnumeration, must be 0
	err := binary.Read(d.r, binary.LittleEndian, &h.RecordTypeEnumeration)
	if err != nil {
		return Header{}, err
	}

	if h.RecordTypeEnumeration != 0 {
		return Header{}, fmt.Errorf("invalid RecordTypeEnumeration, should be 0, got %d", h.RecordTypeEnumeration)
	}

	// ObjectId
	err = binary.Read(d.r, binary.LittleEndian, &h.ObjectId)
	if err != nil {
		return Header{}, err
	}

	// HeaderId
	err = binary.Read(d.r, binary.LittleEndian, &h.HeaderId)
	if err != nil {
		return Header{}, err
	}

	// must be either -1 or 0? not 100% sure, depends on whatever, commented in the meantime
	//if h.HeaderId != -1 && h.HeaderId != 0 {
	//	return Header{}, fmt.Errorf("invalid HeaderId, should be either -1 or 0, got %d", h.HeaderId)
	//}

	// MajorVersion, must be 1
	err = binary.Read(d.r, binary.LittleEndian, &h.MajorVersion)
	if err != nil {
		return Header{}, err
	}

	if h.MajorVersion != 1 {
		return Header{}, fmt.Errorf("invalid MajorVersion, should be 1, got %d", h.MajorVersion)
	}

	// MinorVersion, must be 0
	err = binary.Read(d.r, binary.LittleEndian, &h.MinorVersion)
	if err != nil {
		return Header{}, err
	}

	if h.MinorVersion != 0 {
		return Header{}, fmt.Errorf("invalid MinorVersion, should be 0, got %d", h.MinorVersion)
	}

	return h, nil
}
