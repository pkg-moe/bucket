package s3

func (o *S3) GetCDNDomain() string {
	return o.config.DomainCDN
}
