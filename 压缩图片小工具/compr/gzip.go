package compr

import "bytes"
import "compress/gzip"

type Gzip struct{}

func (gz *Gzip) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write(data)
	if err != nil {
		return nil, err
	}
	err = zw.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (gz *Gzip) Name() string {
	return "gzip"
}
