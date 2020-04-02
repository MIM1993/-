package compr

import (
	"bytes"
	"compress/zlib"
)

type Zlib struct{}

func (bz *Zlib) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
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

func (gz *Zlib) Name() string {
	return "zlib"
}
