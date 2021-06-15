package config

import "os"

const (
	apiGithubAccessToken = "SECRET_github_token"
	LogLevel = "info"
	goEnvironment = "GO_ENVIRONMENT"
	production = "production"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}

func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}