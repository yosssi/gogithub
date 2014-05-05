package gogithub

// consts definition
const (
	NormalURL             = "https://github.com"
	APIURL                = "https://api.github.com"
	AccessTokenURL        = NormalURL + "/login/oauth/access_token"
	AuthenticatedUserURL  = APIURL + "/user"
	URLParamPrefix        = "?"
	ParamKeyAccessToken   = "access_token"
	GetContentsURL        = APIURL + "/repos/%s/%s/contents/%s"
	GetContentURL         = "https://raw.githubusercontent.com/%s/%s/%s/%s"
	SearchRepositoriesURL = APIURL + "/search/repositories"
)
