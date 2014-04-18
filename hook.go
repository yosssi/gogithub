package gogithub

// Hook represents a GitHub hook.
type Hook struct {
	Commits    []Commit   `json:"commits"`
	HeadCommit Commit     `json:"head_commit"`
	Ref        string     `json:"ref"`
	Repository Repository `json:"repository"`
}
