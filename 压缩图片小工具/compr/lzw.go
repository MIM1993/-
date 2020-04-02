package compr

import "bytes"
import "compress/lzw"

type Lzw struct{}

func (gz *Lzw) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := lzw.NewWriter(&buf, lzw.LSB, 8)
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

func (gz *Lzw) Name() string {
	return "lzw"
}
