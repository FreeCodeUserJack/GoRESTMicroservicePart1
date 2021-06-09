package config

import "os"

const (
	apiGithubAccessToken = "SECRET_github_token"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}