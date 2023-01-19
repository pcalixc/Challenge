package models

type QueryTerm struct {
	Term string `json:"term"`
}

type QuerySearchRequest struct {
	SearchType string    `json:"search_type"`
	Query      QueryTerm `json:"query"`
	From       int       `json:"from"`
	MaxResult  int       `json:"max_results"`
	Source     []string  `json:"_source"`
}
