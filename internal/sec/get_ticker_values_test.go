package sec

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetTickerValues(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		if req.URL.String() !=  "/some/path" {
			t.Error("url error: ", req.URL.String())
		}
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))

	c := &client{
		host: "https://sec.gov",
		c:    *server.Client(),
	}

	_, err := c.GetTickerValues()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	// Close the server when test finishes
	defer server.Close()
}