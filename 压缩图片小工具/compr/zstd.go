package compr

import (
	"bytes"
	"github.com/klauspost/compress/zstd"
)

type Zstd struct{}

func (bz *Zstd) Compress(data []byte) ([]byte, error) {

	var buf bytes.Buffer
	zw, err := zstd.NewWriter(&buf, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
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

func (gz *Zstd) Name() string {
	return "zstd"
}
