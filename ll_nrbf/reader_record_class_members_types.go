package ll_nrbf

type RecordClassWithMembersAndTypes struct {
	ClassInfo      interface{}
	MemberTypeInfo interface{}
	LibraryId      int32
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/847b0b6a-86af-4203-8ed0-f84345f845b9
func (d *Decoder) decodeRecordWithMembersAndTypes() (rcmt RecordClassWithMembersAndTypes, err error) {
	rcmt.ClassInfo, err = d.decodeClassInfo()
	if err != nil {
		return
	}

	return
}
