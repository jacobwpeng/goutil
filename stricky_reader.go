package goutil

import "io"

type StrickyReader struct {
	Err error
	r   io.Reader
}

func NewStrickyReader(r io.Reader) *StrickyReader {
	return &StrickyReader{r: r}
}

func (r *StrickyReader) Read(data []byte) (n int, err error) {
	if r.Err != nil {
		return 0, r.Err
	}
	n, r.Err = r.r.Read(data)
	return n, r.Err
}
