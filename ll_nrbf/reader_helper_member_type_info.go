package ll_nrbf

type MemberTypeInfo struct {
	BinaryTypeEnums []byte
	AdditionalInfos []AdditionalInfos
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/aa509b5a-620a-4592-a5d8-7e9613e0a03e
func (d *Decoder) decodeMemberTypeInfo(n int32) (mti MemberTypeInfo, err error) {
	// BinaryTypeEnums
	mti.BinaryTypeEnums = make([]byte, n)
	_, err = d.r.Read(mti.BinaryTypeEnums)
	if err != nil {
		return
	}

	// AdditionalInfos
	mti.AdditionalInfos = make([]AdditionalInfos, n)
	for i := range mti.AdditionalInfos {
		mti.AdditionalInfos[i], err = d.decodeAdditionalInfos(n)
		if err != nil {
			return
		}
	}

	return
}
