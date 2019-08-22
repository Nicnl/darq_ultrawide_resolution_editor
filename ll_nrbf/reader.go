package ll_nrbf

import (
	"io"
)

type Decoder struct {
	r   io.Reader
	buf []byte
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}
