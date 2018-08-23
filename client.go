package client

import (
	rc "github.com/rancher/go-rancher/v3"
	"net/http"
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

	httpClient *http.Client

	projectID  string
	namePrefix string
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

	cli.httpClient = http.DefaultClient
	var project *rc.Project
	projects, err := cli.rancherClient.Project.List(&rc.ListOpts{})
	if err != nil {
		return nil, err
	}

	for _, v := range projects.Data {
		if v.Name == cli.config.ProjectName {
			project = &v
		}
	}

	if project == nil {
		return nil, NotFound
	}
	cli.projectID = project.Id
	cli.SetNamePrefix()

	return cli, nil
}
