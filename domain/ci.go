package domain

// CI represents data from the CI
type CI struct {
	API string `json:"api" graphql:"api"`
	URL string `json:"url" graphql:"url"`
}
