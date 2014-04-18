package gogithub

type Content struct {
	Links   Links   `json:"_links"`
	GitURL  string  `json:"git_url"`
	HtmlURL string  `json:"html_url"`
	Name    string  `json:"name"`
	Path    string  `json:"path"`
	Sha     string  `json:"sha"`
	Size    float64 `json:"size"`
	Type    string  `json:"type"`
	URL     string  `json:"url"`
}
