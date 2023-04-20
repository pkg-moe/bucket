package oss

import (
	"bytes"
	"io"
)

func (o *OSS) Get(name string) ([]byte, error) {
	body, err := o.bucket.GetObject(name)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	data, err := io.ReadAll(body)
	return data, err
}

func (o *OSS) Put(name string, value []byte) error {
	return o.bucket.PutObject(name, bytes.NewReader(value))
}
