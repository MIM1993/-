package compr

import "bytes"
import "github.com/ulikunitz/xz/lzma"

type Lzma struct{}

func (gz *Lzma) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw, err := lzma.NewWriter(&buf)
	if err != nil {
		return nil, err
	}
	_, err = zw.Write(data)
	if err != nil {
		return nil, err
	}
	err = zw.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (gz *Lzma) Name() string {
	return "lzma"
}
