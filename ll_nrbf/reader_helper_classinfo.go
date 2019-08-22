package ll_nrbf

import (
	"encoding/binary"
	"fmt"
	"nrbf/length_prefixed_string"
)

type ClassInfo struct {
	ObjectId    int32
	Name        string
	MemberCount int32
	MemberNames []string
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/0a192be0-58a1-41d0-8a54-9c91db0ab7bf
func (d *Decoder) decodeClassInfo() (ci ClassInfo, err error) {
	// ObjectId
	err = binary.Read(d.r, binary.LittleEndian, &ci.ObjectId)
	if err != nil {
		return
	}

	// Name
	ci.Name, err = length_prefixed_string.Read(d.r)
	if err != nil {
		return
	}

	// MemberCount, 'The value MUST be 0 or a positive integer'
	err = binary.Read(d.r, binary.LittleEndian, &ci.MemberCount)
	if err != nil {
		return
	}

	if ci.MemberCount < 0 {
		err = fmt.Errorf("MemberCount must be 0 or a positive integer, got %d", ci.MemberCount)
		return
	}

	// MemberNames
	ci.MemberNames = make([]string, ci.MemberCount, ci.MemberCount)
	for i := range ci.MemberNames {
		ci.MemberNames[i], err = length_prefixed_string.Read(d.r)
		if err != nil {
			return
		}
	}

	return
}
