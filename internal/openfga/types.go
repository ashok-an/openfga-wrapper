package openfga

type MissingEnv struct {
	Name string
}

type Client struct {
	Url     string
	StoreID string
	ModelID string
}

type Stores struct {
	ContinuationToken string  `json:"continuation_token"`
	Stores            []Store `json:"stores"`
}

type Store struct {
	CreatedAt string  `json:"created_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	UpdatedAt string  `json:"updated_at"`
}

type Model struct {
	ID              string           `json:"id"`
	SchemaVersion   string           `json:"schema_version"`
	TypeDefinitions []TypeDefinition `json:"type_definitions"`
	// Use interface{} for fields with arbitrary nesting
	Conditions interface{} `json:"conditions,omitempty"`
	Metadata   interface{} `json:"metadata,omitempty"`
}

type ModelResponse struct {
	Authorization map[string]Model `json:"authorization_model"`
}

type Models struct {
	ContinuationToken string  `json:"continuation_token"`
	Models            []Model `json:"authorization_models"`
}
type TypeDefinition struct {
	Type string `json:"type"`
	// Use interface{} for fields with arbitrary nesting
	Relations interface{} `json:"relations,omitempty"`
	Metadata  interface{} `json:"metadata,omitempty"`
}
