package compr

import "bytes"
import "github.com/dsnet/compress/bzip2"

type Bzip2 struct{}

func (gz *Bzip2) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	wc := &bzip2.WriterConfig{
		Level: bzip2.BestCompression,
	}
	zw, err := bzip2.NewWriter(&buf, wc)
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

func (gz *Bzip2) Name() string {
	return "bzip2"
}
