package cos

func (c *COS) GetCDNDomain() string {
	return c.config.DomainCDN
}
