package github

type GithubErrorResponse struct {
	StatusCode       int           `json:"status_code"`
	Message          string        `json:"message"`
	Errors           []GithubError `json:"errors"`
	DocumentationUrl string        `json:"documentation_url"`
}

func (g GithubErrorResponse) Error() string {
	return g.Message
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}