package biohelper

import (
	"context"

	"net/http"

	"golang.org/x/oauth2"
)

const token = ""

// GetAuthenticationClient Get authentication context
func GetAuthenticationClient() *http.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	return oauth2.NewClient(ctx, ts)
}
