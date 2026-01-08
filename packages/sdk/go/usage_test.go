// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package strukcode_test

import (
	"context"
	"os"
	"testing"

	"github.com/harutk7/strukcode-sdk-go"
	"github.com/harutk7/strukcode-sdk-go/internal/testutil"
	"github.com/harutk7/strukcode-sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := strukcode.NewClient(
		option.WithBaseURL(baseURL),
	)
	sessions, err := client.Session.List(context.TODO(), strukcode.SessionListParams{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", sessions)
}
