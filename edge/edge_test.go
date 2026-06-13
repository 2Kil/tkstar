package tkEdge

import (
	"net/url"
	"testing"

	"github.com/chromedp/cdproto/network"
)

func setupStore() {
	store.mu.Lock()
	store.latestReqHead = network.Headers{
		"Content-Type": "application/json",
		"X-Custom":     "test-value",
	}
	store.latestResHead = network.Headers{
		"Content-Type": "text/html",
		"X-Response":   "resp-value",
	}
	store.latestQueryParams = url.Values{
		"key": []string{"abc"},
		"id":  []string{"123"},
	}
	store.latestReqID = network.RequestID("req-001")
	store.mu.Unlock()
}

func clearStore() {
	store.mu.Lock()
	store.latestReqHead = nil
	store.latestResHead = nil
	store.latestQueryParams = nil
	store.latestReqID = ""
	store.mu.Unlock()
}

func TestGetResReturnsHeaderValue(t *testing.T) {
	setupStore()
	defer clearStore()

	if got := GetRes("X-Response"); got != "resp-value" {
		t.Fatalf("GetRes(X-Response) = %q, want resp-value", got)
	}
}

func TestGetResReturnsEmptyForMissingKey(t *testing.T) {
	setupStore()
	defer clearStore()

	if got := GetRes("NonExistent"); got != "" {
		t.Fatalf("GetRes(NonExistent) = %q, want empty", got)
	}
}

func TestGetReqReturnsHeaderValue(t *testing.T) {
	setupStore()
	defer clearStore()

	if got := GetReq("Content-Type"); got != "application/json" {
		t.Fatalf("GetReq(Content-Type) = %q, want application/json", got)
	}
}

func TestGetReqCaseInsensitive(t *testing.T) {
	setupStore()
	defer clearStore()

	if got := GetReq("content-type"); got != "application/json" {
		t.Fatalf("GetReq(content-type) = %q, want application/json", got)
	}
}

func TestGetReqReturnsEmptyForMissingKey(t *testing.T) {
	setupStore()
	defer clearStore()

	if got := GetReq("NonExistent"); got != "" {
		t.Fatalf("GetReq(NonExistent) = %q, want empty", got)
	}
}

func TestGetUrlQueryReturnsParamValue(t *testing.T) {
	setupStore()
	defer clearStore()

	if got := GetUrlQuery("key"); got != "abc" {
		t.Fatalf("GetUrlQuery(key) = %q, want abc", got)
	}
	if got := GetUrlQuery("id"); got != "123" {
		t.Fatalf("GetUrlQuery(id) = %q, want 123", got)
	}
}

func TestGetUrlQueryReturnsEmptyForMissingParam(t *testing.T) {
	setupStore()
	defer clearStore()

	if got := GetUrlQuery("missing"); got != "" {
		t.Fatalf("GetUrlQuery(missing) = %q, want empty", got)
	}
}

func TestClearResetsAllStoredData(t *testing.T) {
	setupStore()
	Clear()

	if got := GetRes("X-Response"); got != "" {
		t.Fatalf("after Clear, GetRes = %q, want empty", got)
	}
	if got := GetReq("Content-Type"); got != "" {
		t.Fatalf("after Clear, GetReq = %q, want empty", got)
	}
	if got := GetUrlQuery("key"); got != "" {
		t.Fatalf("after Clear, GetUrlQuery = %q, want empty", got)
	}
}

func TestLoadUrlDoesNotBlockWhenBrowserNotRunning(t *testing.T) {
	// LoadUrl should not panic or block when browser is not running
	LoadUrl("https://example.com")
}

func TestGetUrlReturnsErrorWhenBrowserNotStarted(t *testing.T) {
	browserMu.Lock()
	savedCtx := browserCtx
	browserCtx = nil
	browserMu.Unlock()
	defer func() {
		browserMu.Lock()
		browserCtx = savedCtx
		browserMu.Unlock()
	}()

	_, err := GetUrl()
	if err == nil {
		t.Fatal("GetUrl() expected error when browser not started")
	}
}

func TestGetCookiesReturnsErrorWhenBrowserNotStarted(t *testing.T) {
	browserMu.Lock()
	savedCtx := browserCtx
	browserCtx = nil
	browserMu.Unlock()
	defer func() {
		browserMu.Lock()
		browserCtx = savedCtx
		browserMu.Unlock()
	}()

	_, err := GetCookies()
	if err == nil {
		t.Fatal("GetCookies() expected error when browser not started")
	}
}

func TestGetCookiesAllReturnsErrorWhenBrowserNotStarted(t *testing.T) {
	browserMu.Lock()
	savedCtx := browserCtx
	browserCtx = nil
	browserMu.Unlock()
	defer func() {
		browserMu.Lock()
		browserCtx = savedCtx
		browserMu.Unlock()
	}()

	_, err := GetCookiesAll()
	if err == nil {
		t.Fatal("GetCookiesAll() expected error when browser not started")
	}
}

func TestStopDoesNotPanicWhenBrowserNotRunning(t *testing.T) {
	Stop()
}
