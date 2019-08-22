package ll_nrbf

import (
	"fmt"
	"io"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) WriteRecord(rec Record) (err error) {
	_, err = e.w.Write([]byte{byte(rec.RecordType)})
	if err != nil {
		return err
	}

	switch rec.RecordType {
	case RTE_0_SERIALIZED_STREAM_HEADER:
		err = e.encodeSerializedStreamHeader(rec.Record.(SerializedStreamHeader))
	default:
		err = fmt.Errorf("record type %d not implemented", rec.RecordType)
	}

	return
}
