package protocol

import (
	"bytes"
	"compress/zlib"
	"io"
)

func zlibParser(data []byte) []byte {
	buf := bytes.NewReader(data)
	reader, err := zlib.NewReader(buf)
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal("zlib解压错误 " + err.Error())
	}
	return b
}

func brotliParser(data []byte) []byte {
	log.Fatal("function not defined")
	return nil
}
