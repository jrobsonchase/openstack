package v2_0

import (
	//"net/http"
	"net/url"
	"errors"
	//"regexp"
	"github.com/Pursuit92/openstack/core"
)

var (
	ErrNoCreds error = errors.New("No credentials specified.")
	ErrNoTenant = errors.New("No tenant specified.")
	ErrNotAuthed = errors.New("Not authenticated.")
)

type IdentityClient struct {
	// The token for authentication. Id is the only required field for
	// authentication purposes.
	token *Token
	passwordCredentials *PasswordCredentials
	tenantName string
	tenantId string
	AuthUrl string
	Access *Access
}

// Creates a new Identity client. Will return any errors from parsing
// the url, but performs no network operations
func NewClient(authUrl string) (*IdentityClient,error) {
	url,err := url.Parse(authUrl)
	if err != nil {
		return nil,err
	}

	var client IdentityClient
	client.AuthUrl = url.String()

	return &client,nil
}

// Authenticates the Identity client and stores the results of the attempt
// in the Access field. One of Token or PasswordCredentials is required to be set.
// The token value is used first if both exist.
// One of TenantName or TenantId is required to be set.
// TenantId is used first if both exist.
func (ic *IdentityClient) Authenticate() error {
	if nil == ic.token && nil == ic.passwordCredentials {
		return ErrNoCreds
	}

	var authInfo Auth
	if ic.token != nil {
		authInfo = Auth{Token: ic.token}
	} else {
		authInfo = Auth{PasswordCredentials: ic.passwordCredentials}
	}

	if ic.tenantName == "" && ic.tenantId == "" {
		return ErrNoTenant
	}

	if ic.tenantId != "" {
		authInfo.TenantId = ic.tenantId
	} else {
		authInfo.TenantName = ic.tenantName
	}

	resp := &tokensResponse{}
	err := core.OsRequest("POST",ic.AuthUrl + "/tokens",tokensRequest{&authInfo},resp,"")
	if err != nil {
		return err
	}
	ic.Access = resp.Access

	return nil
}

func (ic *IdentityClient) PasswordAuth(user,pass string) {
	ic.passwordCredentials = &PasswordCredentials{user,pass}
}

func (ic *IdentityClient) TokenAuth(tokenStr string) {
	ic.token = &Token{Id: tokenStr}
}

func (ic *IdentityClient) TenantName(tn string) {
	ic.tenantName = tn
}

func (ic *IdentityClient) TenantId(ti string) {
	ic.tenantId = ti
}

func (ic *IdentityClient) AuthedReq(method, url string, data, resp interface{}) error {
	if ic.Access == nil {
		return ErrNotAuthed
	}

	return core.OsRequest(method,url,data,resp,ic.Access.Token.Id)
}

