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

func TestFindFilesWithOptionalParams(t *testing.T) {
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
	_, err := client.Find.Files(context.TODO(), strukcode.FindFilesParams{
		Query:     strukcode.F("query"),
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

func TestFindSymbolsWithOptionalParams(t *testing.T) {
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
	_, err := client.Find.Symbols(context.TODO(), strukcode.FindSymbolsParams{
		Query:     strukcode.F("query"),
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

func TestFindTextWithOptionalParams(t *testing.T) {
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
	_, err := client.Find.Text(context.TODO(), strukcode.FindTextParams{
		Pattern:   strukcode.F("pattern"),
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
