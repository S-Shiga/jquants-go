//go:build fulltest

package jquants

import (
	"context"
	"testing"
)

func TestClient_IssueInformation(t *testing.T) {
	ctx := context.Background()
	if err := setup(ctx); err != nil {
		t.Fatalf("Failed to setup client: %v", err)
	}
	resp, err := testClient.IssueInformation(ctx, IssueInformationRequest{})
	if err != nil {
		t.Errorf("Failed to get issue information: %v", err)
	}
	if len(resp) == 0 {
		t.Error("Empty response")
	}
}
