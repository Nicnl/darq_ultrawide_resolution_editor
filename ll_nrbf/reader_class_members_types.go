package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

type RecordClassWithMembersAndTypes struct {
	ClassInfo      ClassInfo
	MemberTypeInfo MemberTypeInfo
	LibraryId      int32
	Values         []interface{}
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
	rcmt.Values = make([]interface{}, rcmt.ClassInfo.MemberCount)
	for i := range rcmt.Values {
		fmt.Println("## DOING i =", i)
		fmt.Println("## TYPE IS:", rcmt.MemberTypeInfo.BinaryTypeEnums[i])
		switch rcmt.MemberTypeInfo.BinaryTypeEnums[i] {
		case BTE_0_PRIMITIVE:
			rcmt.Values[i], err = d.decodePrimitive(rcmt.MemberTypeInfo.AdditionalInfos[i].Data.(PrimitiveType))
		case BTE_1_STRING:
			var nextByte uint8
			nextByte, err = d.nextByte()
			if err != nil {
				return
			}

			if nextByte != 0x06 {
				err = fmt.Errorf("invalid nextByte for type BTE_1_STRING, expected 6, got %d", nextByte)
				return
			}

			rcmt.Values[i], err = d.decodeBinaryObjectString()
		default:
			err = fmt.Errorf("decoding not implemented for type %d", rcmt.MemberTypeInfo.BinaryTypeEnums[i])
		}

		if err != nil {
			return
		}

		fmt.Printf("rcmt.Values[%d] = %s\n", i, fmt.Sprint(rcmt.Values[i]))
	}

	return
}
