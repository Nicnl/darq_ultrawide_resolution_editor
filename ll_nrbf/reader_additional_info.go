package ll_nrbf

import "fmt"

type AdditionalInfo struct {
	Valid bool
	Data  interface{}
}

//3 4

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/aa509b5a-620a-4592-a5d8-7e9613e0a03e
func (d *Decoder) decodeAdditionalInfo(bte BinaryType, n int32) (ai AdditionalInfo, err error) {
	// 'The AdditionalInfos sequence MUST NOT contain any item for the BinaryTypeEnum values of String, Object, ObjectArray, or StringArray.'
	if bte == BTE_1_STRING || bte == BTE_2_OBJECT || bte == BTE_5_OBJECT_ARRAY || bte == BTE_6_STRING_ARRAY {
		ai.Valid = false
		ai.Data = nil
		return
	}

	if bte == BTE_0_PRIMITIVE || bte == BTE_7_PRIMITIVE_ARRAY {
		// PrimitiveTypeEnumeration
		ai.Valid = true
		ai.Data, err = d.decodePrimitiveType()
		return
	}

	if bte == BTE_3_SYSTEM_CLASS {
		// String
		err = fmt.Errorf("AdditionalInfo not implemented for ClassTypeInfo")
		return
	}

	if bte == BTE_4_CLASS {
		// ClassTypeInfo
		err = fmt.Errorf("AdditionalInfo not implemented for ClassTypeInfo")
		return
	}

	err = fmt.Errorf("wtf, this is not supposed to happen (invalid BinaryTypeEnumeration supplied)")
	return
}
