package v2

import (
	keystone "github.com/Pursuit92/openstack/identity/v2_0"
	"errors"
	//"github.com/Pursuit92/openstack/core"
)

var (
	ErrNoComputeSvc = errors.New("No compute service in catalog.")
)

type ComputeClient struct {
	*keystone.IdentityClient
	Endpoint keystone.Endpoint
}

func NewClient(authUrl string) (*ComputeClient, error) {
	cc := &ComputeClient{}
	var err error
	cc.IdentityClient,err = keystone.NewClient(authUrl)
	return cc,err
}

func (cc *ComputeClient) Authenticate() error {
	err := cc.IdentityClient.Authenticate()
	if err != nil {
		return err
	}

	found := false
	for _,v := range  cc.Access.ServiceCatalog {
		if v.Type == "compute" {
			if len(v.Endpoints) >= 1 {
				cc.Endpoint = v.Endpoints[0]
				found = true
				break
			}
		}
	}
	if ! found {
		return ErrNoComputeSvc
	}
	return nil
}

func (cc *ComputeClient) Servers() ([]*Server,error) {
	resp := make(map[string][]*Server)
	err := cc.AuthedReq("GET",cc.Endpoint.PublicUrl + "/servers",nil,&resp)
	if err != nil {
		return nil,err
	}
	return resp["servers"],nil
}

func (cc *ComputeClient) ServersDetail() ([]*Server,error) {
	resp := make(map[string][]*Server)
	err := cc.AuthedReq("GET",cc.Endpoint.PublicUrl + "/servers/detail",nil,&resp)
	if err != nil {
		return nil,err
	}
	return resp["servers"],err
}

func (cc *ComputeClient) Details(serverId string) (*Server,error) {
	resp := make(map[string]*Server)

	err := cc.AuthedReq("GET",cc.Endpoint.PublicUrl + "/servers/" + serverId,nil,&resp)
	if err != nil {
		return nil,err
	}
	return resp["server"],nil
}

func (cc *ComputeClient) Delete(srv *Server) error {
	serverId := srv.Id
	return cc.AuthedReq("DELETE",cc.Endpoint.PublicUrl + "/servers/" + serverId,nil,nil)
}

func (cc *ComputeClient) Update(srv *Server) (*Server,error) {
	resp := make(map[string]*Server)
	err := cc.AuthedReq("PUT",cc.Endpoint.PublicUrl + "/servers/" + srv.Id,srv,&resp)
	if err != nil {
		return nil,err
	}
	return resp["server"],nil
}

func NewServer() *Server {
	return &Server{ServerDetail: &ServerDetail{},ServerCreate: &ServerCreate{}}
}

func (cc *ComputeClient) Create(srv *Server) (*Server,error) {
	req := make(map[string]*Server)
	resp := make(map[string]*Server)
	if srv.TenantId == "" {
		srv.TenantId = cc.Access.Token.Tenant.Id
	}
	req["server"] = srv
	err := cc.AuthedReq("POST",cc.Endpoint.PublicUrl + "/servers",req, &resp)
	if err != nil {
		return nil,err
	}
	return resp["server"],nil
}