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

## Prerequisite: Install go-swagger

Before generating code, ensure you have the correct tap for go-swagger:

```sh
brew tap go-swagger/go-swagger
brew install go-swagger
```