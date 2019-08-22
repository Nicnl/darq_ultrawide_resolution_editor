package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

type BinaryTypeEnumeration uint8

const (
	BTE_0_PRIMITIVE       = 0
	BTE_1_STRING          = 1
	BTE_2_OBJECT          = 2
	BTE_3_SYSTEM_CLASS    = 3
	BTE_4_CLASS           = 4
	BTE_5_OBJECT_ARRAY    = 5
	BTE_6_STRING_ARRAY    = 6
	BTE_7_PRIMITIVE_ARRAY = 7
)

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/054e5c58-be21-4c86-b1c3-f6d3ce17ec72
func (d *Decoder) decodeBinaryTypeEnumeration() (bte BinaryTypeEnumeration, err error) {
	err = binary.Read(d.r, binary.LittleEndian, &bte)
	if err != nil {
		return
	}

	if bte != BTE_0_PRIMITIVE &&
		bte != BTE_1_STRING &&
		bte != BTE_2_OBJECT &&
		bte != BTE_3_SYSTEM_CLASS &&
		bte != BTE_4_CLASS &&
		bte != BTE_5_OBJECT_ARRAY &&
		bte != BTE_6_STRING_ARRAY &&
		bte != BTE_7_PRIMITIVE_ARRAY {
		err = fmt.Errorf("invalid BinaryTypeEnumeration, expected one of [%d, %d, %d, %d, %d, %d, %d, %d], got %d",
			BTE_0_PRIMITIVE,
			BTE_1_STRING,
			BTE_2_OBJECT,
			BTE_3_SYSTEM_CLASS,
			BTE_4_CLASS,
			BTE_5_OBJECT_ARRAY,
			BTE_6_STRING_ARRAY,
			BTE_7_PRIMITIVE_ARRAY,
			bte,
		)
	}

	return
}
