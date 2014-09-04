package core

import ()

type Link struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
}

type OsObj struct {
	Name  string `json:"name,omitempty"`
	Id    string `json:"id,omitempty"`
	Links []Link `json:"links,omitempty"`
}
