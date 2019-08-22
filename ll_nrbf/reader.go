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
