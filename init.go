package bucket

import (
	"github.com/pkg/errors"

	"pkg.moe/pkg/bucket/dao"
)

func NewBucket(typeName Type, config *bucketdao.Config) (bucketdao.IBucket, error) {
	b := typeToNewBucket(typeName)
	if b == nil {
		return nil, errors.New("error bucket type")
	}
	err := b.Init(config)
	return b, err
}
