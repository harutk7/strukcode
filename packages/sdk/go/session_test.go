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

func TestSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.New(context.TODO(), strukcode.SessionNewParams{
		Directory: strukcode.F("directory"),
		ParentID:  strukcode.F("sesJ!"),
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

func TestSessionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Update(
		context.TODO(),
		"id",
		strukcode.SessionUpdateParams{
			Directory: strukcode.F("directory"),
			Title:     strukcode.F("title"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.List(context.TODO(), strukcode.SessionListParams{
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

func TestSessionDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Delete(
		context.TODO(),
		"sesJ!",
		strukcode.SessionDeleteParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionAbortWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Abort(
		context.TODO(),
		"id",
		strukcode.SessionAbortParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionChildrenWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Children(
		context.TODO(),
		"sesJ!",
		strukcode.SessionChildrenParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionCommandWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Command(
		context.TODO(),
		"id",
		strukcode.SessionCommandParams{
			Arguments: strukcode.F("arguments"),
			Command:   strukcode.F("command"),
			Directory: strukcode.F("directory"),
			Agent:     strukcode.F("agent"),
			MessageID: strukcode.F("msgJ!"),
			Model:     strukcode.F("model"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Get(
		context.TODO(),
		"sesJ!",
		strukcode.SessionGetParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionInitWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Init(
		context.TODO(),
		"id",
		strukcode.SessionInitParams{
			MessageID:  strukcode.F("msgJ!"),
			ModelID:    strukcode.F("modelID"),
			ProviderID: strukcode.F("providerID"),
			Directory:  strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Message(
		context.TODO(),
		"id",
		"messageID",
		strukcode.SessionMessageParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Messages(
		context.TODO(),
		"id",
		strukcode.SessionMessagesParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Prompt(
		context.TODO(),
		"id",
		strukcode.SessionPromptParams{
			Parts: strukcode.F([]strukcode.SessionPromptParamsPartUnion{strukcode.TextPartInputParam{
				Text: strukcode.F("text"),
				Type: strukcode.F(strukcode.TextPartInputTypeText),
				ID:   strukcode.F("id"),
				Metadata: strukcode.F(map[string]interface{}{
					"foo": "bar",
				}),
				Synthetic: strukcode.F(true),
				Time: strukcode.F(strukcode.TextPartInputTimeParam{
					Start: strukcode.F(0.000000),
					End:   strukcode.F(0.000000),
				}),
			}}),
			Directory: strukcode.F("directory"),
			Agent:     strukcode.F("agent"),
			MessageID: strukcode.F("msgJ!"),
			Model: strukcode.F(strukcode.SessionPromptParamsModel{
				ModelID:    strukcode.F("modelID"),
				ProviderID: strukcode.F("providerID"),
			}),
			NoReply: strukcode.F(true),
			System:  strukcode.F("system"),
			Tools: strukcode.F(map[string]bool{
				"foo": true,
			}),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionRevertWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Revert(
		context.TODO(),
		"id",
		strukcode.SessionRevertParams{
			MessageID: strukcode.F("msgJ!"),
			Directory: strukcode.F("directory"),
			PartID:    strukcode.F("prtJ!"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionShareWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Share(
		context.TODO(),
		"id",
		strukcode.SessionShareParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionShellWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Shell(
		context.TODO(),
		"id",
		strukcode.SessionShellParams{
			Agent:     strukcode.F("agent"),
			Command:   strukcode.F("command"),
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionSummarizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Summarize(
		context.TODO(),
		"id",
		strukcode.SessionSummarizeParams{
			ModelID:    strukcode.F("modelID"),
			ProviderID: strukcode.F("providerID"),
			Directory:  strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUnrevertWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Unrevert(
		context.TODO(),
		"id",
		strukcode.SessionUnrevertParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUnshareWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Unshare(
		context.TODO(),
		"sesJ!",
		strukcode.SessionUnshareParams{
			Directory: strukcode.F("directory"),
		},
	)
	if err != nil {
		var apierr *strukcode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
