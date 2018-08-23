package client

import (
	rc "github.com/rancher/go-rancher/v3"
)

type Config struct {
	Url         string
	ProjectName string
	AccessKey   string
	SecretKey   string
}

type Client struct {
	config        *Config
	rancherClient *rc.RancherClient
}

func NewClient(config *Config) (*Client, error) {
	cli := &Client{
		config: config,
	}

	rancherClient, err := rc.NewRancherClient(&rc.ClientOpts{
		Url:       config.Url,
		AccessKey: config.AccessKey,
		SecretKey: config.SecretKey,
	})

	if err != nil {
		return nil, err
	}

	cli.rancherClient = rancherClient

	return cli, nil
}
