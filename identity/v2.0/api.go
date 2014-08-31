package api

import (
	//"net/http"
	"net/url"
	//"regexp"
)

type IdentityClient struct {
	Token *token
	PasswordCredentials *passwordCredentials
	TenantName string
	TenantId string
	AuthUrl *url.URL
	Access *access
	User *user
	Metadata *metadata
}

func NewClient(authUrl string) (*IdentityClient,error) {
	var client IdentityClient
	var err error
	client.AuthUrl,err = url.Parse(authUrl)
	if err != nil {
		return nil,err
	}
	return &client,nil
}

