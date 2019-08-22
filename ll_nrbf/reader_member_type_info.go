package ll_nrbf

type MemberTypeInfo struct {
	BinaryTypeEnums []BinaryTypeEnumeration
	AdditionalInfos []AdditionalInfo
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/aa509b5a-620a-4592-a5d8-7e9613e0a03e
func (d *Decoder) decodeMemberTypeInfo(n int32) (mti MemberTypeInfo, err error) {
	// BinaryTypeEnums
	mti.BinaryTypeEnums = make([]BinaryTypeEnumeration, n)
	for i := range mti.BinaryTypeEnums {
		mti.BinaryTypeEnums[i], err = d.decodeBinaryTypeEnumeration()
		if err != nil {
			return
		}
	}

	// AdditionalInfos
	mti.AdditionalInfos = make([]AdditionalInfo, n)
	for i := range mti.AdditionalInfos {
		mti.AdditionalInfos[i], err = d.decodeAdditionalInfo(mti.BinaryTypeEnums[i], n)
		if err != nil {
			return
		}
	}

	return
}
