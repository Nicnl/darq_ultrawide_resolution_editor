package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

type RecordClassWithMembersAndTypes struct {
	ClassInfo      ClassInfo
	MemberTypeInfo MemberTypeInfo
	LibraryId      int32
	Values         []Record
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/847b0b6a-86af-4203-8ed0-f84345f845b9
func (d *Decoder) decodeRecordWithMembersAndTypes() (rcmt RecordClassWithMembersAndTypes, err error) {
	// ClassInfo
	rcmt.ClassInfo, err = d.decodeClassInfo()
	if err != nil {
		return
	}

	// MemberTypeInfo
	rcmt.MemberTypeInfo, err = d.decodeMemberTypeInfo(rcmt.ClassInfo.MemberCount)
	if err != nil {
		return
	}

	// LibraryId
	err = binary.Read(d.r, binary.LittleEndian, &rcmt.LibraryId)
	if err != nil {
		return
	}

	// Values
	rcmt.Values = make([]Record, rcmt.ClassInfo.MemberCount)
	for i := range rcmt.Values {
		rcmt.Values[i], err = d.NextRecord()
		if err != nil {
			return
		}
		fmt.Printf("rcmt.Values[%d] = %s\n", i, fmt.Sprint(rcmt.Values[i]))
	}

	return
}
