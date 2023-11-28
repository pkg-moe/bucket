package ucloud_us3

import (
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"

	"pkg.moe/pkg/bucket/dao"
)

type UCloudUS3 struct {
	config *bucketdao.Config

	clientConfig *ufsdk.Config
}

func (o *UCloudUS3) Init(config *bucketdao.Config) (err error) {
	o.config = config

	o.clientConfig = &ufsdk.Config{
		PublicKey:       config.AK,
		PrivateKey:      config.SK,
		BucketName:      config.Bucket,
		FileHost:        config.Endpoint,
		Endpoint:        config.DomainCDN,
		VerifyUploadMD5: true,
	}
	return
}
