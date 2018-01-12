package client

import (
	vault "github.com/hashicorp/vault/api"
	gentlemen "gopkg.in/h2non/gentleman.v2"
)

//Client interface
type Client interface {
	Write(path string, data map[string]interface{}, token string) (*vault.Secret, error)
	Read(path string, token string) (*vault.Secret, error)
	CreatePolicy(path string, role string, policyName string) error
	CreateToken(policyName string) (string, error)
}

type client struct {
	token      string
	httpclient *gentlemen.Client
}

//NewClient creates a new client with the corresponding Vault token and corresponding path to Vault
func NewClient(token string, path string) (Client, error) {
	var httpClient = gentlemen.New()
	httpClient.BaseURL(path)
	return &client{
		token:      token,
		httpclient: httpClient,
	}, nil
}
