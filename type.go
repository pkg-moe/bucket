package bucket

import (
	"pkg.moe/pkg/bucket/cos"
	"pkg.moe/pkg/bucket/dao"
	"pkg.moe/pkg/bucket/oss"
)

type Type string

const (
	OSS Type = "oss"
	COS Type = "cos"
)

func typeToNewBucket(typeName Type) bucketdao.IBucket {
	switch typeName {
	case OSS:
		return &oss.OSS{}
	case COS:
		return &cos.COS{}
	}

	return nil
}
