package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

type ClassWithMembersAndTypes struct {
	ClassInfo      ClassInfo
	MemberTypeInfo MemberTypeInfo
	LibraryId      int32
	Values         []interface{}
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/847b0b6a-86af-4203-8ed0-f84345f845b9
func (d *Decoder) decodeClassWithMembersAndTypes() (cmt ClassWithMembersAndTypes, err error) {
	// ClassInfo
	cmt.ClassInfo, err = d.decodeClassInfo()
	if err != nil {
		return
	}

	// MemberTypeInfo
	cmt.MemberTypeInfo, err = d.decodeMemberTypeInfo(cmt.ClassInfo.MemberCount)
	if err != nil {
		return
	}

	// LibraryId
	err = binary.Read(d.r, binary.LittleEndian, &cmt.LibraryId)
	if err != nil {
		return
	}

	// Values
	cmt.Values = make([]interface{}, cmt.ClassInfo.MemberCount)
	for i := range cmt.Values {
		switch cmt.MemberTypeInfo.BinaryTypeEnums[i] {
		case BTE_0_PRIMITIVE:
			cmt.Values[i], err = d.decodePrimitive(cmt.MemberTypeInfo.AdditionalInfos[i].Data.(PrimitiveType))
		case BTE_1_STRING:
			cmt.Values[i], err = d.NextRecord()
		default:
			err = fmt.Errorf("decoding not implemented for type %d", cmt.MemberTypeInfo.BinaryTypeEnums[i])
		}

		if err != nil {
			return
		}
	}

	return
}
