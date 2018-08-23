package client

import (
	"encoding/json"
	"errors"
	rc "github.com/rancher/go-rancher/v3"
	"io/ioutil"
	"net/http"
)

var (
	NotFound = errors.New("not found")
	ReqFail  = errors.New("req fail")
)

type Cert struct {
	Name  string `json:"name"`
	Key   string `json:"key"`
	Certs string `json:"certs"`
}

func (c Cert) ToCertificate() rc.Certificate {
	return rc.Certificate{
		Name: c.Name,
		Key:  c.Key,
		Cert: c.Certs,
	}
}

func (c *Client) CertGet(cert Cert) (*rc.Certificate, error) {
	path := c.getCertResPathByName(cert.Name)
	resp, err := c.doReq(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, NotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ReqFail
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var cer rc.Certificate
	err = json.Unmarshal(body, &cer)
	return &cer, nil
}

func (c *Client) CertUpdate(cert Cert) error {
	path := c.getCertResPathByName(cert.Name)
	resp, err := c.doReq(http.MethodPut, path, cert)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return ReqFail
	}
	return nil
}

func (c *Client) CertAdd(cert Cert) error {
	path := c.getCertResPath()
	resp, err := c.doReq(http.MethodPost, path, cert)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return ReqFail
	}
	return nil
}
