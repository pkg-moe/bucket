package bucketdao

type Config struct {
	DomainCDN string

	Endpoint    string
	Bucket      string
	AK          string
	SK          string
	Region      string
	PartitionID string
}

type ObjectInfo struct {
	Key          string
	LastModified int64
}

type IBucket interface {
	// init
	Init(config *Config) (err error)

	// base
	Get(name string) ([]byte, error)
	Put(name string, value []byte) error

	// sign
	SignGetUrl(name string, expiredInSec int64) (string, error)
	SignPutUrl(name string, expiredInSec int64) (string, error)
	SignPutUrlWithMD5(name string, expiredInSec int64, md5 string) (string, error)

	// info
	GetCDNDomain() string

	//List
	ListObjectsV2(path string) ([]ObjectInfo, error)
}
