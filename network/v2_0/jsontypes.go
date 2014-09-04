package v2_0

type Network struct {
	Bridge            string `json:"bridge,omitempty"`
	BridgeInterface   string `json:"bridge_interface,omitempty"`
	Broadcast         string `json:"broadcast,omitempty"`
	Cidr              string `json:"cidr,omitempty"`
	Cidr6             string `json:"cidr_v6,omitempty"`
	CreatedAt         string `json:"created_at,omitempty"`
	Deleted           bool   `json:"deleted,omitempty"`
	DeletedAt         string `json:"deleted_at,omitempty"`
	DhcpStart         string `json:"dhcp_start,omitempty"`
	Dns1              string `json:"dns1,omitempty"`
	Dns2              string `json:"dns2,omitempty"`
	Gateway           string `json:"gateway,omitempty"`
	Gateway6          string `json:"gateway_v6,omitempty"`
	Host              string `json:"host,omitempty"`
	Id                string `json:"id,omitempty"`
	Injected          bool   `json:"injected,omitempty"`
	Label             string `json:"label,omitempty"`
	MultiHost         bool   `json:"multi_host,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	Netmask6          string `json:"netmask_v6,omitempty"`
	Priority          string `json:"priority,omitempty"`
	ProjectId         string `json:"project_id,omitempty"`
	RxtxBase          string `json:"rxtx_base,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	Vlan              int    `json:"vlan,omitempty"`
	VpnPrivateAddress string `json:"vpn_private_address,omitempty"`
	VpnPublicAddress  string `json:"vpn_public_address,omitempty"`
	VpnPublicPort     int    `json:"vpn_public_port,omitempty"`
}
