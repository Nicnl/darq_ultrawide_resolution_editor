package ll_nrbf

import (
	"encoding/binary"
	"fmt"
	"nrbf/length_prefixed_string"
)

type BinaryLibrary struct {
	LibraryId   int32
	LibraryName string
}

// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-nrbf/7fcf30e1-4ad4-4410-8f1a-901a4a1ea832
// [0C] 02 00 00 00 0F 41 73 73 65 6D 62 6C 79 2D 43 53 68 61 72 70 05 01 00 00 00 19 4D 61 69 6E 47 61 6D 65 53 61 76 65 44 61 74 61 2B 53 61 76 65 44 61 74 61 2D 00 00 00 0C 67 61 6D 65 4C 61 6E 67 75 61 67 65 0E 67 61 6D 65 4C 61 6E 67 75
func (d *Decoder) decodeRecordBinaryLibrary() (bl BinaryLibrary, err error) {
	// LibraryId, must be positive
	err = binary.Read(d.r, binary.LittleEndian, &bl.LibraryId)
	if err != nil {
		return BinaryLibrary{}, err
	}

	if bl.LibraryId < 0 {
		return BinaryLibrary{}, fmt.Errorf("invalid LibraryId, should be positive, got %d", bl.LibraryId)
	}

	// LibraryName
	libraryName, err := length_prefixed_string.Read(d.r)
	if err != nil {
		return BinaryLibrary{}, err
	}
	bl.LibraryName = libraryName

	return bl, nil
}
