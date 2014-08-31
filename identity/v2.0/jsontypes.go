package api

type passwordCredentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type token struct {
	Id string `json:"id,omitempty"`
	IssuedAt string `json:"issued_at,omitempty"`
	Expires string `json:"expires,omitempty"`
	Tenant *tenant `json:"tenant,omitempty"`
}

type tenant struct {
	Description string `json:"description,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type auth struct {
	TenantName string `json:"tenantName,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	PasswordCredentials *passwordCredentials `json:"passwordCredentials,omitempty"`
	Token *token `json:"token,omitempty"`
}

type tokensRequest struct {
	TenantName string `json:"tenantName,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	Auth *auth `json:"auth,omitempty"`
}

type access struct {
	Token *token `json:"token,omitempty"`
}

type catalogEntry struct {
	Endpoints []endpoint `json:"endpoints,omitempty"`
	EndpointsLinks []string `json:"endpoints_links,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

type endpoint struct {
	AdminURL string `json:"adminURL,omitempty"`
	Region string `json:"region,omitempty"`
	InternalURL string `json:"internalURL,omitempty"`
	Id string `json:"id,omitempty"`
	PublicUrl string `json:"publicURL,omitempty"`
}

type user struct {
	Username string `json:"username,omitempty"`
	RolesLinks []string `json:"roles_links,omitempty"`
	Id string `json:"id,omitempty"`
	Roles []role `json:"roles,omitempty"`
	Name string `json:"name,omitempty"`
}

type role struct {
	Name string `json:"name,omitempty"`
}

type metadata struct {
	IsAdmin int `json:"is_admin,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type tokensResponse struct {
	Access *access `json:"access,omitempty"`
	ServiceCatalog []catalogEntry `json:"serviceCatalog,omitempty"`
	User *user `json:"user,omitempty"`
	Metadata *metadata `json:"metadata,omitempty"`
}
