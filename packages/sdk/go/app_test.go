// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package strukcode_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/harutk7/strukcode-sdk-go"
	"github.com/harutk7/strukcode-sdk-go/internal/testutil"
	"github.com/harutk7/strukcode-sdk-go/option"
)

func TestAppLogWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.App.Log(context.TODO(), strukcode.AppLogParams{
		Level:     strukcode.F(strukcode.AppLogParamsLevelDebug),
		Message:   strukcode.F("message"),
		Service:   strukcode.F("service"),
		Directory: strukcode.F("directory"),
		Extra: strukcode.F(map[string]interface{}{
			"foo": "bar",
		}),
	})
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppProvidersWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.App.Providers(context.TODO(), strukcode.AppProvidersParams{
		Directory: strukcode.F("directory"),
	})
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
