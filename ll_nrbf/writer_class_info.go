package ll_nrbf

import (
	"encoding/binary"
	"fmt"
	"nrbf/length_prefixed_string"
)

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/0a192be0-58a1-41d0-8a54-9c91db0ab7bf
func (e *Encoder) encodeClassInfo(ci ClassInfo) (err error) {
	// ObjectId
	err = binary.Write(e.w, binary.LittleEndian, ci.ObjectId)
	if err != nil {
		return
	}

	// Name
	err = length_prefixed_string.Write(e.w, ci.Name)
	if err != nil {
		return
	}

	// MemberCount, 'The value MUST be 0 or a positive integer'
	if ci.MemberCount < 0 {
		err = fmt.Errorf("MemberCount must be 0 or a positive integer, got %d", ci.MemberCount)
		return
	}

	err = binary.Write(e.w, binary.LittleEndian, ci.MemberCount)
	if err != nil {
		return
	}

	// MemberNames
	ci.MemberNames = make([]string, ci.MemberCount, ci.MemberCount)
	for i := range ci.MemberNames {
		err = length_prefixed_string.Write(e.w, ci.MemberNames[i])
		if err != nil {
			return
		}
	}

	return
}
