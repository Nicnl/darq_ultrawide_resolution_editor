package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

func (e *Encoder) encodePrimitiveType(pte PrimitiveType) (err error) {
	if pte != PTE_1_BOOLEAN &&
		pte != PTE_2_BYTE &&
		pte != PTE_3_CHAR &&
		pte != PTE_5_DECIMAL &&
		pte != PTE_6_DOUBLE &&
		pte != PTE_7_INT16 &&
		pte != PTE_8_INT32 &&
		pte != PTE_9_INT64 &&
		pte != PTE_10_SBYTE &&
		pte != PTE_11_SINGLE &&
		pte != PTE_12_TIME_SPAN &&
		pte != PTE_13_DATE_TIME &&
		pte != PTE_14_UINT16 &&
		pte != PTE_15_UINT32 &&
		pte != PTE_16_UINT64 &&
		pte != PTE_17_NULL &&
		pte != PTE_18 {
		err = fmt.Errorf("invalid BinaryTypeEnumeration, expected one of [%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d], got %d",
			PTE_1_BOOLEAN,
			PTE_2_BYTE,
			PTE_3_CHAR,
			PTE_5_DECIMAL,
			PTE_6_DOUBLE,
			PTE_7_INT16,
			PTE_8_INT32,
			PTE_9_INT64,
			PTE_10_SBYTE,
			PTE_11_SINGLE,
			PTE_12_TIME_SPAN,
			PTE_13_DATE_TIME,
			PTE_14_UINT16,
			PTE_15_UINT32,
			PTE_16_UINT64,
			PTE_17_NULL,
			PTE_18,
			pte,
		)
	}

	err = binary.Write(e.w, binary.LittleEndian, pte)
	if err != nil {
		return
	}

	return
}
