package cos

import (
	"bytes"
	"context"
	"io"
)

func (c *COS) Get(name string) ([]byte, error) {
	resp, err := c.client.Object.Get(context.Background(), name, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, err
}

func (c *COS) Put(name string, value []byte) error {
	_, err := c.client.Object.Put(context.Background(), name, bytes.NewReader(value), nil)
	return err
}
