package bucket

import (
	"pkg.moe/pkg/bucket/cos"
	"pkg.moe/pkg/bucket/dao"
	"pkg.moe/pkg/bucket/oss"
	"pkg.moe/pkg/bucket/s3"
	"pkg.moe/pkg/bucket/ucloud_us3"
)

type Type string

const (
	OSS        Type = "oss"
	COS        Type = "cos"
	S3         Type = "s3"
	UCLOUD_US3 Type = "ucloud_us3"
)

func typeToNewBucket(typeName Type) bucketdao.IBucket {
	switch typeName {
	case OSS:
		return &oss.OSS{}
	case COS:
		return &cos.COS{}
	case S3:
		return &s3.S3{}
	case UCLOUD_US3:
		return &ucloud_us3.UCloudUS3{}
	}

	return nil
}
