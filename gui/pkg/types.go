package pkg

type Rule struct {
	Policies []Policy `json:"policies"`
}

type Policy struct {
	UriResource string      `json:"uri_resource"`
	Action      []string    `json:"action"`
	Statements  []Statement `json:"statements"`
}

type Statement struct {
	Type string   `json:"type"`
	Uri  string   `json:"uri"`
	Func FuncData `json:"func"`
}

type FuncData struct {
	Type string   `json:"type"`
	Uri  string   `json:"uri"`
	Term TermData `json:"term"`
}

type TermData struct {
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

type PublishRequest struct {
	Name string `json:"name"`
	Rule *Rule  `json:"rule"`
}

type TemplateResponse struct {
	Rule   *Rule   `json:"rule"`
	Render *string `json:"policy"`
}

type Settings struct {
	NameBundle  string         `json:"name"`
	Repository  string         `json:"repository"`
	GithubToken string         `json:"token"`
	AzureCrend  *ACredentials  `json:"crendentials"`
	TypeBundle  string         `json:"type"`
	Data        map[string]any `json:"data"`
}

type ACredentials struct {
	StorageAccountName string `json:"storageAccountName"`
	TenantID           string `json:"tenantID"`
	ClientID           string `json:"clientId"`
	ClientSecret       string `json:"clientSecret"`
	ContainerName      string `json:"containerName"`
}

type Manifest struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
