package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/054e5c58-be21-4c86-b1c3-f6d3ce17ec72
func (e *Encoder) encodeBinaryType(bte BinaryType) (err error) {
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

	err = binary.Write(e.w, binary.LittleEndian, bte)
	if err != nil {
		return
	}

	return
}
