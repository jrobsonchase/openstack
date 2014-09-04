package v2

import (
	"github.com/Pursuit92/openstack/core"
)

type Image struct {
	Id    string      `json:"id,omitempty"`
	Links []core.Link `json:"links,omitempty"`
}
