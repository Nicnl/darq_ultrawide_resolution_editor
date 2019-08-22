package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

type PrimitiveTypeEnumeration uint8

const (
	PTE_1_BOOLEAN    = 1
	PTE_2_BYTE       = 2
	PTE_3_CHAR       = 3
	PTE_5_DECIMAL    = 5
	PTE_6_DOUBLE     = 6
	PTE_7_INT16      = 7
	PTE_8_INT32      = 8
	PTE_9_INT64      = 9
	PTE_10_SBYTE     = 10
	PTE_11_SINGLE    = 11
	PTE_12_TIME_SPAN = 12
	PTE_13_DATE_TIME = 13
	PTE_14_UINT16    = 14
	PTE_15_UINT32    = 15
	PTE_16_UINT64    = 16
	PTE_17_NULL      = 17
	PTE_18           = 18
)

func (d *Decoder) decodePrimitiveTypeEnumeration() (pte PrimitiveTypeEnumeration, err error) {
	err = binary.Read(d.r, binary.LittleEndian, &pte)
	if err != nil {
		return
	}

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

	return
}
