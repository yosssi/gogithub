package gogithub

const (
	NormalURL            = "https://github.com"
	APIURL               = "https://api.github.com"
	AccessTokenURL       = NormalURL + "/login/oauth/access_token"
	AuthenticatedUserURL = APIURL + "/user"
	URLParamPrefix       = "?"
	ParamKeyAccessToken  = "access_token"
	GetContentsPath      = APIURL + "/repos/%s/%s/contents/%s"
	GetContentPath       = "https://raw.githubusercontent.com/%s/%s/%s/%s"
)
