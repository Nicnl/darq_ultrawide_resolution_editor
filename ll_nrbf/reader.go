package ll_nrbf

import (
	"io"
)

type Decoder struct {
	r   io.Reader
	buf []byte
}

// https://stackoverflow.com/questions/3052202/how-to-analyse-contents-of-binary-serialization-stream/30176566#30176566
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) nextByte() (b byte, err error) {
	var p [1]byte
	_, err = d.r.Read(p[:])
	if err == nil {
		b = p[0]
	}
	return
}
