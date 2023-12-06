package ucloud_us3

import (
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
)

func (o *UCloudUS3) DeleteObjects(keys []string) (err error) {
	req, err := ufsdk.NewFileRequest(o.clientConfig, nil)
	if err != nil {
		return err
	}

	for _, key := range keys {
		if err := req.DeleteFile(key); err != nil {
			return err
		}
	}

	return
}
