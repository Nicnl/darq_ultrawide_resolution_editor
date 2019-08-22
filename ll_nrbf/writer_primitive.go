package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

func (e *Encoder) encodePrimitive(pte PrimitiveType, value interface{}) (err error) {
	switch pte {
	case PTE_1_BOOLEAN:
		v := uint8(0)
		if value.(bool) { // TODO: fix risky cast
			v = 1
		}

		err = binary.Write(e.w, binary.LittleEndian, v)
	case PTE_2_BYTE:
		err = binary.Write(e.w, binary.LittleEndian, value.(uint8)) // TODO: fix risky cast
	case PTE_3_CHAR:
		err = binary.Write(e.w, binary.LittleEndian, value.(byte)) // TODO: fix risky cast
	case PTE_6_DOUBLE:
		err = binary.Write(e.w, binary.LittleEndian, value.(float64)) // TODO: fix risky cast
	case PTE_7_INT16:
		err = binary.Write(e.w, binary.LittleEndian, value.(int16)) // TODO: fix risky cast
	case PTE_8_INT32:
		err = binary.Write(e.w, binary.LittleEndian, value.(int32)) // TODO: fix risky cast
	case PTE_9_INT64:
		err = binary.Write(e.w, binary.LittleEndian, value.(int64)) // TODO: fix risky cast
	case PTE_11_SINGLE:
		err = binary.Write(e.w, binary.LittleEndian, value.(float32)) // TODO: fix risky cast
	case PTE_14_UINT16:
		err = binary.Write(e.w, binary.LittleEndian, value.(uint16)) // TODO: fix risky cast
	case PTE_15_UINT32:
		err = binary.Write(e.w, binary.LittleEndian, value.(uint32)) // TODO: fix risky cast
	case PTE_16_UINT64:
		err = binary.Write(e.w, binary.LittleEndian, value.(uint64)) // TODO: fix risky cast
	default:
		err = fmt.Errorf("primitive type not implemented: %d", pte)
	}

	if err != nil {
		return
	}

	return
}
