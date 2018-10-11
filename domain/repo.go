package domain

// Repo holds information on a repository
type Repo struct {
	// Configurable fields
	URL  string `json:"url" graphql:"url" mapstructure:"url"`
	Path string `json:"path" graphql:"path" mapstructure:"path"`
	Type string `json:"type" graphql:"type" mapstructure:"type"`
	Host string `json:"host" graphql:"host" mapstructure:"host"`

	// API filled fields
	Description string `json:"description" graphql:"description"`
	Stars       int64  `json:"stars" graphql:"stars"`
	Forks       int64  `json:"forks" graphql:"forks"`
	Watchers    int64  `json:"watchers" graphql:"watchers"`
}
