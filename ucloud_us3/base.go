package ucloud_us3

import (
	"bytes"
	"io"
	"time"

	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
)

func (o *UCloudUS3) Get(name string) ([]byte, error) {
	req, err := ufsdk.NewFileRequest(o.clientConfig, nil)
	if err != nil {
		return nil, err
	}

	reqUrl := req.GetPrivateURL(name, 10*time.Minute)

	buf := bytes.NewBuffer(nil)
	if err := req.DownloadFile(buf, reqUrl); err != nil {
		return nil, err
	}

	data, err := io.ReadAll(buf)
	return data, err
}

func (o *UCloudUS3) Put(name string, value []byte) error {
	req, err := ufsdk.NewFileRequest(o.clientConfig, nil)
	if err != nil {
		return err
	}

	return req.IOPut(bytes.NewReader(value), name, "")
}
