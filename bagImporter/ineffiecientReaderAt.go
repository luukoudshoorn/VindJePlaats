package bagImporter

import (
	"io"
	"io/ioutil"
)

// Solution from https://stackoverflow.com/questions/40245442/handling-nested-zip-files-with-archive-zip
type inefficientReaderAt struct {
	rdr    io.ReadCloser
	cur    int64
	initer func() (io.ReadCloser, error)
}

func newInefficentReaderAt(initer func() (io.ReadCloser, error)) *inefficientReaderAt {
	return &inefficientReaderAt{
		initer: initer,
	}
}

func (r *inefficientReaderAt) Close() error {
	return r.rdr.Close()
}

func (r *inefficientReaderAt) Read(p []byte) (n int, err error) {
	n, err = r.rdr.Read(p)
	if err == io.EOF && n > 0 {
		// reader was wrong to return EOF, fix that here
		err = nil
	}
	r.cur += int64(n)
	return n, err
}

func (r *inefficientReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	// reset on rewind
	if off < r.cur || r.rdr == nil {
		r.cur = 0
		r.rdr, err = r.initer()
		if err != nil {
			return 0, err
		}
	}

	if off > r.cur {
		sz, err := io.CopyN(ioutil.Discard, r.rdr, off-r.cur)
		n = int(sz)
		r.cur = off

		if err != nil {
			return n, err
		}
	}

	return r.Read(p)
}
