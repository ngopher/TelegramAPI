package mtprotox
import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"io"
	"log"
)
const data="02930293"
func do() {
	raw, err := hex.DecodeString(data)
	if err != nil {
		log.Fatal(err)
	}

	data, err := gunzip(raw)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Decompressed: %x", data)
}

func gunzip(compressed []byte) ([]byte, error) {
	decompressor, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, decompressor)
	if err != nil {
		return nil, err
	}

	err = decompressor.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
