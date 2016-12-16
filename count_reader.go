package goutil

import "io"

type CountReader struct {
	r io.Reader
	n int64
}

func NewCountReader(r io.Reader) *CountReader {
	return &CountReader{r: r}
}

func (r *CountReader) Count() int64 {
	return r.n
}

func (r *CountReader) Read(data []byte) (n int, err error) {
	n, err = r.r.Read(data)
	r.n += int64(n)
	return n, err
}
