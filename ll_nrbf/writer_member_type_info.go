package ll_nrbf

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/aa509b5a-620a-4592-a5d8-7e9613e0a03e
func (e *Encoder) encodeMemberTypeInfo(mti MemberTypeInfo) (err error) {
	// BinaryTypeEnums
	for i := range mti.BinaryTypeEnums {
		err = e.encodeBinaryType(mti.BinaryTypeEnums[i])
		if err != nil {
			return
		}
	}

	// AdditionalInfos
	for i := range mti.AdditionalInfos {
		err = e.encodeAdditionalInfo(mti.BinaryTypeEnums[i], mti.AdditionalInfos[i])
		if err != nil {
			return
		}
	}

	return
}
