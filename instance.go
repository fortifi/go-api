//go:generate $GOPATH/src/github.com/fortifi/go-api/generate.sh

package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fortifi/go-api/client/authentication"

	"github.com/fortifi/go-api/client"
	"github.com/fortifi/go-api/models"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	expiryBuffer             = 14400
	fortifiOrgHeader         = "x-fortifi-org"
	fortifiAuthHeader        = "authorization"
	fortifiBearerTokenSchema = "Bearer %s"
)

var (
	debugFortifiRequests = false
	unsecureSchemes      = []string{"http"}
)

// Instance is a single org API instance
type Instance struct {
	expiry              int64
	apiHost             string
	organisationFID     string
	authToken           string
	user                string
	key                 string
	apiInstance         *client.FortifiAPI
	authenticator       *Authenticator
	hostingProductGroup string
	domainProductGroup  string
}

// Authenticator auths fortifi requests
type Authenticator struct {
	organisationFID string
	authtoken       string
}

// NewInstance creates a new FortifiInstance with embeded tokenhelper
func NewInstance(orgFid, user, key string) (Instance, error) {
	inst := Instance{}
	inst.SetOrganisationFID(orgFid)
	inst.SetUser(user)
	inst.SetKey(key)
	return inst, nil
}

// AuthenticateRequest handles request authentication
func (a *Authenticator) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	req.SetHeaderParam(fortifiOrgHeader, a.organisationFID)
	req.SetHeaderParam(fortifiAuthHeader, fmt.Sprintf(fortifiBearerTokenSchema, a.authtoken))
	return nil
}

func (a *Authenticator) AuthenticateStdRequest(req http.Request) error {
	req.Header.Set(fortifiOrgHeader, a.organisationFID)
	req.Header.Set(fortifiAuthHeader, fmt.Sprintf(fortifiBearerTokenSchema, a.authtoken))
	return nil
}

// InitAPI called by main function to start the fortifi API and webhook listener
func (f *Instance) InitAPI() error {
	_, err := f.GetAPIInstance()
	return err
}

// GetAPIInstance returns current instance of fortifi API
func (f *Instance) GetAPIInstance() (*client.FortifiAPI, error) {
	if f.apiInstance != nil && time.Now().Unix() < (f.expiry-expiryBuffer) {
		return f.apiInstance, nil
	}

	if f.organisationFID == "" {
		return nil, errors.New("Organization FID is required but not set")
	}

	if f.user == "" {
		return nil, errors.New("User is required but not set")
	}

	if f.key == "" {
		return nil, errors.New("Key is required but not set")
	}

	transport := httptransport.New(f.GetAPIHost(), client.DefaultBasePath, f.GetSchemes())
	transport.DefaultAuthentication = f.GetAuthenticator()
	transport.SetDebug(debugFortifiRequests)
	err := f.getNewToken(transport)
	if err != nil {
		return nil, fmt.Errorf("Failed to authenticate with Fortifi: %s", err.Error())
	}

	f.apiInstance = client.New(NewTransport(transport, f.getNewToken), strfmt.Default)
	return f.apiInstance, nil
}

func (f *Instance) getNewToken(transport runtime.ClientTransport) error {
	c := client.New(transport, strfmt.Default)
	p := authentication.NewGetServiceAuthTokenParams()
	p.Payload = &models.ServiceAccountCredentialsPayload{
		ID:  &f.user,
		Key: &f.key,
	}

	re, err := c.Authentication.GetServiceAuthToken(p)
	if err != nil {
		return err
	}

	f.expiry = re.Payload.Data.Expiry
	f.authToken = re.Payload.Data.Token
	return nil
}

// GetAuthenticator returns fortifi authenticator instance
func (f *Instance) GetAuthenticator() *Authenticator {
	return &Authenticator{organisationFID: f.organisationFID, authtoken: f.authToken}
}

// GetSchemes returns the API protocol
func (f *Instance) GetSchemes() []string {
	if strings.HasPrefix(f.apiHost, "http://") {
		return unsecureSchemes
	}
	return client.DefaultSchemes
}

// SetLocalDebug determines if local debug vars and logging are used for API
func (f *Instance) SetLocalDebug(s bool) {
	debugFortifiRequests = s
}

// SetAPIHost sets fortifi API endpoint
func (f *Instance) SetAPIHost(host string) {
	f.apiHost = host
}

// GetAPIHost returns the API host (no schema as per swagger format)
func (f *Instance) GetAPIHost() string {
	host := client.DefaultHost
	if f.apiHost != "" {
		host = strings.TrimPrefix(f.apiHost, "https://")
		host = strings.TrimPrefix(host, "http://")
	}
	return host
}

// SetOrganisationFID sets the organisation fid used by the fortifi API
// this is case sensitive
func (f *Instance) SetOrganisationFID(of string) {
	f.organisationFID = of
}

// SetUser sets the user ID (This should be a service account) used by the fortifi API
func (f *Instance) SetUser(u string) {
	f.user = u
}

// SetKey sets the key for given User ID used by the fortifi API
func (f *Instance) SetKey(k string) {
	f.key = k
}
