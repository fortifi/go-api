# Fortifi go-api

## Usage
Import api with your required client & models
```go
import (
	api "github.com/fortifi/go-api"
	"github.com/fortifi/go-api/client/customers"
)
```

New Fortifi api instance with your service account credentials
```go
	fortifi, err := api.NewInstance(
		"my-org-fid",
		"my-service-account-user",
		"my-service-account-key")
```

Use Fortifi
```go
	// Example: Retrieve customer by external reference
	cid := "my-customer-external-reference"

	params := customers.NewGetCustomersFindByReferenceParams()
	params.SetReference(&cid)

	instance, err := fortifi.GetAPIInstance()
	if err != nil {
		panic(fmt.Sprintf("API instance not available %s\n", err.Error()))
	}

	response, err := instance.Customers.GetCustomersFindByReference(params, fortifi.GetAuthenticator())
	if err != nil {
		panic(fmt.Sprintf("failed to retrieve customer because %s\n", err.Error()))
	}

	fmt.Printf("Customers First Name is %s\n", response.Payload.FirstName)

```

## Disable nesting for a call

Pass `api.WithDisableNesting` as the final argument to a client method:

```go
response, err := instance.Customers.GetCustomersFindByReference(
	params,
	fortifi.GetAuthenticator(),
	api.WithDisableNesting,
)
```

This sends `Disable-Nesting: true` for that call, asking the server to skip
automatic expansion of related FIDs. It also works with the `...Context`
methods. Calls without this option retain the server's default behaviour,
even when reusing the same parameters. The endpoint must honour the header;
Citadel procedures may control nesting independently on their backend calls.

To request compact related objects containing `id`, `fid`, and `displayName`,
use `api.WithShortNesting` instead:

```go
response, err := instance.Customers.GetCustomersFindByReference(
	params,
	fortifi.GetAuthenticator(),
	api.WithShortNesting,
)
```

This sends `Short-Nesting: true` for that call and also supports `...Context`
methods. If both options are supplied, the backend gives `Disable-Nesting`
precedence.

## Prerequisite: Install go-swagger

Before generating code, ensure you have the correct tap for go-swagger:

```sh
brew tap go-swagger/go-swagger
brew install go-swagger
```
