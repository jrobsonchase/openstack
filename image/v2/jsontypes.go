package v2

import (
	"github.com/Pursuit92/openstack/core"
)

type Image struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
	Status string `json:"status,omitempty"`
	Progress int `json:"progress,omitempty"`
	MinDisk int `json:"minDisk,omitempty"`
	MinRam int `json:"minRam,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Links []core.Link `json:"links,omitempty"`
}
