package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

func (d *Decoder) decodePrimitive(pte PrimitiveType) (out interface{}, err error) {
	switch pte {
	case PTE_1_BOOLEAN:
		var v uint8
		err = binary.Read(d.r, binary.LittleEndian, &v)
		fmt.Println("v =", v)
		if err == nil {
			out = v != 0
		}
	case PTE_2_BYTE:
		var v uint8
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_3_CHAR:
		var v byte
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_6_DOUBLE:
		var v float64
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_7_INT16:
		var v int16
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_8_INT32:
		var v int32
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_9_INT64:
		var v int64
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_11_SINGLE:
		var v float32
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_14_UINT16:
		var v uint16
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_15_UINT32:
		var v uint32
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	case PTE_16_UINT64:
		var v uint64
		err = binary.Read(d.r, binary.LittleEndian, &v)
		if err == nil {
			out = v
		}
	default:
		err = fmt.Errorf("primitive type not implemented: %d", pte)
	}

	if err != nil {
		return
	}

	return
}
