package gogithub

const (
	// NormalURL represents a GitHub normal URL.
	NormalURL = "https://github.com"
	// APIURL represents a GitHub API URL.
	APIURL = "https://api.github.com"
	// AccessTokenURL represents an OAuth access token API URL.
	AccessTokenURL = NormalURL + "/login/oauth/access_token"
	// AuthenticatedUserURL represents a get user API URL.
	AuthenticatedUserURL = APIURL + "/user"
	// URLParamPrefix represents a URL parameter prefix.
	URLParamPrefix = "?"
	// ParamKeyAccessToken represents a parameter key of an access token.
	ParamKeyAccessToken = "access_token"
	// GetContentsURL represents a get contents API URL.
	GetContentsURL = APIURL + "/repos/%s/%s/contents/%s"
	// GetContentURL represents a get content API URL.
	GetContentURL = "https://raw.githubusercontent.com/%s/%s/%s/%s"
	// SearchRepositoriesURL represents a search repositories API URL.
	SearchRepositoriesURL = APIURL + "/search/repositories"
)
