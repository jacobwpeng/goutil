package goutil

import "io"

type StrickyWriter struct {
	Err error
	w   io.Writer
}

func NewStrickyWriter(w io.Writer) *StrickyWriter {
	return &StrickyWriter{w: w}
}

func (w *StrickyWriter) Write(data []byte) (n int, err error) {
	if w.Err != nil {
		return 0, w.Err
	}
	n, w.Err = w.w.Write(data)
	return n, w.Err
}
