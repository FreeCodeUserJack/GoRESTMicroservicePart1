package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestConstants(t *testing.T) {
	assert.EqualValues(t, apiGithubAccessToken, "SECRET_github_token")
}

func TestGetGithubAccessToken(t *testing.T) {
	assert.NotEqual(t, "", githubAccessToken)
}