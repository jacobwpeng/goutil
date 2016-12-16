package goutil

import "io"

type CountWriter struct {
	w io.Writer
	n int64
}

func NewCountWriter(w io.Writer) *CountWriter {
	return &CountWriter{w: w}
}

func (w *CountWriter) Count() int64 {
	return w.n
}

func (w *CountWriter) Write(data []byte) (n int, err error) {
	n, err := w.Write(data)
	w.n += int64(n)
	return n, err
}
