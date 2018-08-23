package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) setupReq(req *http.Request) {
	req.SetBasicAuth(c.config.AccessKey, c.config.SecretKey)
}

func (c *Client) SetNamePrefix() {
	arr := strings.Split(c.projectID, ":")
	c.namePrefix = arr[1]
}

func (c *Client) doReq(method, path string, body interface{}) (*http.Response, error) {
	var input io.Reader
	if body != nil {
		s, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		input = strings.NewReader(string(s))
	}

	req, err := http.NewRequest(method, c.getUri(path), input)
	if err != nil {
		return nil, err
	}

	c.setupReq(req)
	res, err := c.httpClient.Do(req)
	return res, err
}

func (c *Client) getUri(path string) string {
	return fmt.Sprintf("%s/%s", c.config.Url, path)
}

func (c *Client) getCertResPath() string {
	return fmt.Sprintf("project/%s/certificates", c.projectID)
}

func (c *Client) getCertResPathByName(name string) string {
	return fmt.Sprintf("%s/%s", c.getCertResPath(), fmt.Sprintf("%s:%s", c.namePrefix, name))
}
