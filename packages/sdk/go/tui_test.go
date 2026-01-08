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

func TestTuiAppendPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.AppendPrompt(context.TODO(), strukcode.TuiAppendPromptParams{
		Text:      strukcode.F("text"),
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

func TestTuiClearPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ClearPrompt(context.TODO(), strukcode.TuiClearPromptParams{
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

func TestTuiExecuteCommandWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ExecuteCommand(context.TODO(), strukcode.TuiExecuteCommandParams{
		Command:   strukcode.F("command"),
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

func TestTuiOpenHelpWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenHelp(context.TODO(), strukcode.TuiOpenHelpParams{
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

func TestTuiOpenModelsWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenModels(context.TODO(), strukcode.TuiOpenModelsParams{
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

func TestTuiOpenSessionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenSessions(context.TODO(), strukcode.TuiOpenSessionsParams{
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

func TestTuiOpenThemesWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenThemes(context.TODO(), strukcode.TuiOpenThemesParams{
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

func TestTuiShowToastWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ShowToast(context.TODO(), strukcode.TuiShowToastParams{
		Message:   strukcode.F("message"),
		Variant:   strukcode.F(strukcode.TuiShowToastParamsVariantInfo),
		Directory: strukcode.F("directory"),
		Title:     strukcode.F("title"),
	})
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiSubmitPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.SubmitPrompt(context.TODO(), strukcode.TuiSubmitPromptParams{
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
