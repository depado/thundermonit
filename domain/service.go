package domain

import "time"

// Service is a single service
type Service struct {
	// Configurable fields
	ID   int64  `json:"id" graphql:"id" mapstructure:"id"`
	Name string `json:"name" graphql:"name" mapstructure:"name"`
	URL  string `json:"url" graphql:"url" mapstructure:"url"`
	Host string `json:"host" graphql:"host" mapstructure:"host"`
	Icon string `json:"icon" graphql:"icon" mapstructure:"icon"`
	Own  bool   `json:"own" graphql:"own" mapstructure:"own"`

	// Auto-filled fields
	Latest          string        `json:"latest" graphql:"latest"`
	RespTime        time.Duration `json:"respTime" graphql:"respTime"`
	Status          int           `json:"status" graphql:"status"`
	CurrentBuildURL string        `json:"currentBuildUrl" graphql:"currentBuildUrl"`

	Repo *Repo `json:"repo,omitempty" graphql:"repo"`
	CI   *CI   `json:"ci,omitempty" graphql:"ci"`
}
