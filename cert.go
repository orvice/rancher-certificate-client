package client

import (
	rc "github.com/rancher/go-rancher/v3"
)

type Cert struct {
	Name      string
	Key, Cert string
}

func (c Cert) ToCertificate() rc.Certificate {
	return rc.Certificate{
		Name: c.Name,
		Key:  c.Key,
		Cert: c.Cert,
	}
}
