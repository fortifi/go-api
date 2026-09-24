package api

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/fortifi/go-api/client"
	"github.com/fortifi/go-api/client/customers"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type nestingRoundTripper func(*http.Request) (*http.Response, error)

func (f nestingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestNestingOptions(t *testing.T) {
	for _, tc := range []struct {
		name   string
		header string
		option func(*runtime.ClientOperation)
	}{
		{"disable", "Disable-Nesting", WithDisableNesting},
		{"short", "Short-Nesting", WithShortNesting},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var headers []string
			httpClient := &http.Client{Transport: nestingRoundTripper(func(req *http.Request) (*http.Response, error) {
				headers = append(headers, req.Header.Get(tc.header))
				if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
					t.Errorf("Authorization = %q", got)
				}
				if got := req.Header.Get("X-Fortifi-Org"); got != "test-org" {
					t.Errorf("X-Fortifi-Org = %q", got)
				}
				if got := req.URL.Query().Get("reference"); got != "test-reference" {
					t.Errorf("reference = %q", got)
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     http.Header{"Content-Type": []string{"application/json"}},
					Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
					Request:    req,
				}, nil
			})}
			transport := httptransport.NewWithClient("example.test", "/", []string{"https"}, httpClient)
			auth := &Authenticator{organisationFID: "test-org", authtoken: "test-token"}
			transport.DefaultAuthentication = auth
			c := client.New(NewTransport(transport, nil, nil), strfmt.Default)
			reference := "test-reference"
			params := customers.NewGetCustomersFindByReferenceParams().WithReference(&reference)

			if _, err := c.Customers.GetCustomersFindByReference(params, auth, tc.option); err != nil {
				t.Fatal(err)
			}
			// Reusing params must not carry the header into subsequent calls.
			if _, err := c.Customers.GetCustomersFindByReference(params, auth); err != nil {
				t.Fatal(err)
			}
			// Context methods and transport-provided authentication work too.
			if _, err := c.Customers.GetCustomersFindByReferenceContext(context.Background(), params, nil, tc.option); err != nil {
				t.Fatal(err)
			}
			if len(headers) != 3 || headers[0] != "true" || headers[1] != "" || headers[2] != "true" {
				t.Fatalf("%s headers = %q; want [true, empty, true]", tc.header, headers)
			}
		})
	}
}
