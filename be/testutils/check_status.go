package testutils

import (
	"io"
	"net/http"
	"testing"
)

func AssertResponseStatus(t *testing.T, got *http.Response, status int) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })
	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got.StatusCode != status {
		t.Fatalf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}
}
