package length_prefixed_string

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRead(t *testing.T) {
	input := []byte{
		0x0F, 0x41, 0x73, 0x73, 0x65, 0x6D, 0x62, 0x6C, 0x79, 0x2D, 0x43, 0x53, 0x68, 0x61, 0x72, 0x70, 0x05, 0x01,
		0x00, 0x00, 0x00, 0x19, 0x4D, 0x61, 0x69, 0x6E, 0x47, 0x61, 0x6D, 0x65, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61,
		0x74, 0x61, 0x2B, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2D, 0x00, 0x00, 0x00, 0x0C, 0x67, 0x61,
		0x6D, 0x65, 0x4C, 0x61, 0x6E, 0x67, 0x75, 0x61, 0x67, 0x65, 0x0E, 0x67, 0x61, 0x6D, 0x65, 0x4C,
	}

	res, err := Read(bytes.NewReader(input))
	assert.NoError(t, err)
	assert.Equal(t, "Assembly-CSharp", res)
}
