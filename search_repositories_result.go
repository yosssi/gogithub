package gogithub

// SearchRepositoriesResult represents a search repositories result.
type SearchRepositoriesResult struct {
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
	TotalCount        int64        `json:"total_count"`
}
