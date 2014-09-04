package v2_0

type PasswordCredentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Token struct {
	Id string `json:"id,omitempty"`
	IssuedAt string `json:"issued_at,omitempty"`
	Expires string `json:"expires,omitempty"`
	Tenant *Tenant `json:"tenant,omitempty"`
}

type Tenant struct {
	Description string `json:"description,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Auth struct {
	TenantName string `json:"tenantName,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	PasswordCredentials *PasswordCredentials `json:"passwordCredentials,omitempty"`
	Token *Token `json:"token,omitempty"`
}

type tokensRequest struct {
	Auth *Auth `json:"auth,omitempty"`
}

type Access struct {
	Token *Token `json:"token,omitempty"`
	ServiceCatalog []CatalogEntry `json:"serviceCatalog,omitempty"`
	User *User `json:"user,omitempty"`
	Metadata *metadata `json:"metadata,omitempty"`
}

type CatalogEntry struct {
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	EndpointsLinks []string `json:"endpoints_links,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

type Endpoint struct {
	AdminURL string `json:"adminURL,omitempty"`
	Region string `json:"region,omitempty"`
	InternalURL string `json:"internalURL,omitempty"`
	Id string `json:"id,omitempty"`
	PublicUrl string `json:"publicURL,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`
	RolesLinks []string `json:"roles_links,omitempty"`
	Id string `json:"id,omitempty"`
	Roles []Role `json:"roles,omitempty"`
	Name string `json:"name,omitempty"`
}

type Role struct {
	Name string `json:"name,omitempty"`
}

type metadata struct {
	IsAdmin int `json:"is_admin,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type tokensResponse struct {
	Access *Access `json:"access,omitempty"`
}
