package compr

import (
	"testing"
)

var data []byte = []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

func TestGzip(t *testing.T) {
	var gz Gzip

	res, err := gz.Compress(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("compress rate:%f", float32(len(res))/float32(len(data)))
}
func TestZlib(t *testing.T) {
	var zl Zlib

	res, err := zl.Compress(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("compress rate:%f", float32(len(res))/float32(len(data)))
}

func TestLzw(t *testing.T) {
	var lzw Lzw

	res, err := lzw.Compress(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("compress rate:%f", float32(len(res))/float32(len(data)))
}

func TestBzip2(t *testing.T) {
	var bz Bzip2

	res, err := bz.Compress(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("compress rate:%f", float32(len(res))/float32(len(data)))
}

func TestLzma(t *testing.T) {
	var lz Lzma

	res, err := lz.Compress(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("compress rate:%f", float32(len(res))/float32(len(data)))
}

func TestZstd(t *testing.T) {
	var lz Zstd

	res, err := lz.Compress(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("compress rate:%f", float32(len(res))/float32(len(data)))
}
