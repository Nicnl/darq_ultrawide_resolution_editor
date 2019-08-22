package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/847b0b6a-86af-4203-8ed0-f84345f845b9
func (e *Encoder) encodeClassWithMembersAndTypes(cmt ClassWithMembersAndTypes) (err error) {
	// ClassInfo
	err = e.encodeClassInfo(cmt.ClassInfo)
	if err != nil {
		return
	}

	// MemberTypeInfo
	err = e.encodeMemberTypeInfo(cmt.MemberTypeInfo)
	if err != nil {
		return
	}

	// LibraryId
	err = binary.Write(e.w, binary.LittleEndian, cmt.LibraryId)
	if err != nil {
		return
	}

	// Values
	for i := range cmt.Values {
		switch cmt.MemberTypeInfo.BinaryTypeEnums[i] {
		case BTE_0_PRIMITIVE:
			err = e.encodePrimitive(cmt.MemberTypeInfo.AdditionalInfos[i].Data.(PrimitiveType), cmt.Values[i])
		case BTE_1_STRING:
			err = e.WriteRecord(cmt.Values[i].(Record))
		default:
			err = fmt.Errorf("decoding not implemented for type %d", cmt.MemberTypeInfo.BinaryTypeEnums[i])
		}

		if err != nil {
			return
		}
	}

	return
}
