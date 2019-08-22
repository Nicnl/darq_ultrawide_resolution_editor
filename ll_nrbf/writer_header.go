package ll_nrbf

import (
	"encoding/binary"
	"fmt"
)

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/a7e578d3-400a-4249-9424-7529d10d1b3c
// [00] [01 00 00 00] [FF FF FF FF] [01 00 00 00] [00 00 00 00]
// Respectively: RecordTypeEnumeration, ObjectId, HeaderId, MajorVersion and MinorVersion
func (e *Encoder) encodeSerializedStreamHeader(ssh SerializedStreamHeader) (err error) {
	// ObjectId
	err = binary.Write(e.w, binary.LittleEndian, ssh.ObjectId)
	if err != nil {
		return
	}

	// HeaderId
	err = binary.Write(e.w, binary.LittleEndian, ssh.HeaderId)
	if err != nil {
		return
	}

	// must be either -1 or 0? not 100% sure, depends on whatever, commented in the meantime
	//if ssh.HeaderId != -1 && ssh.HeaderId != 0 {
	//	err = fmt.Errorf("invalid HeaderId, should be either -1 or 0, got %d", ssh.HeaderId)
	//	return
	//}

	// MajorVersion, must be 1
	if ssh.MajorVersion != 1 {
		err = fmt.Errorf("invalid MajorVersion, should be 1, got %d", ssh.MajorVersion)
		return
	}

	err = binary.Write(e.w, binary.LittleEndian, ssh.MajorVersion)
	if err != nil {
		return
	}

	// MinorVersion, must be 0
	if ssh.MinorVersion != 0 {
		err = fmt.Errorf("invalid MinorVersion, should be 0, got %d", ssh.MinorVersion)
		return
	}

	err = binary.Write(e.w, binary.LittleEndian, ssh.MinorVersion)
	if err != nil {
		return
	}

	return
}
