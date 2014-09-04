package v2

import (
	"github.com/Pursuit92/openstack/core"
	glance "github.com/Pursuit92/openstack/image/v2"
)

// Fully describes a server
type Server struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Links []core.Link `json:"links,omitempty"`
	*ServerDetail
	*ServerCreate
	*CreateResp
}

// fields unique to detailed server information
type ServerDetail struct {
	Updated string `json:"updated,omitempty"`
	TenantId string `json:"tenant_id,omitempty"`
	Status string `json:"status,omitempty"`
	Progress int `json:"progress,omitempty"`
	Image glance.Image `json:"image,omitempty"`
	HostId string `json:"host_id,omitempty"`
	Flavor Flavor `json:"flavor,omitempty"`
	Created string `json:"created,omitempty"`
	Addresses map[string][]Address `json:"addresses,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	AccessIPv4 string `json:"accessIPv4,omitempty"`
	AccessIPv6 string `json:"accessIPv6,omitempty"`
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`
	ConfigDrive string `json:"config_drive,omitempty"`
}

type Flavor struct {
	Id string `json:"id,omitempty"`
	Links []core.Link `json:"links,omitempty"`
}

type Address struct {
	Addr string `json:"addr,omitempty"`
	Version int `json:"version,omitempty"`
}

// fields unique to server creation
type ServerCreate struct {
	UserData string `json:"user_data,omitempty"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
	ImageRef string `json:"imageRef,omitempty"`
	FlavorRef string `json:"flavorRef,omitempty"`
	Networks []Network `json:"networks,omitempty"`
	Personality map[string]string `json:"personality,omitempty"`
}

// fields unique to the server creation response
type CreateResp struct {
	AdminPass string `json:"adminPass,omitempty"`
}

type Network struct {
	// necessary if the port isn't specified
	Uuid string `json:"uuid,omitempty"`
	// necessary if the uuid isn't specified
	Port string `json:"port,omitempty"`
	FixedIp string `json:"fixed_ip,omitempty"`
}

type SecurityGroup struct {
	Name string `json:"name,omitempty"`
}
