// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new orders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for orders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteOrdersOrderFidOffersOfferFid removes an offer from an order
*/
func (a *Client) DeleteOrdersOrderFidOffersOfferFid(params *DeleteOrdersOrderFidOffersOfferFidParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrdersOrderFidOffersOfferFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrdersOrderFidOffersOfferFidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOrdersOrderFidOffersOfferFid",
		Method:             "DELETE",
		PathPattern:        "/orders/{orderFid}/offers/{offerFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrdersOrderFidOffersOfferFidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrdersOrderFidOffersOfferFidOK), nil

}

/*
DeleteOrdersOrderFidProductsOrderProductFid removes a product from an order
*/
func (a *Client) DeleteOrdersOrderFidProductsOrderProductFid(params *DeleteOrdersOrderFidProductsOrderProductFidParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrdersOrderFidProductsOrderProductFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrdersOrderFidProductsOrderProductFidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOrdersOrderFidProductsOrderProductFid",
		Method:             "DELETE",
		PathPattern:        "/orders/{orderFid}/products/{orderProductFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrdersOrderFidProductsOrderProductFidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrdersOrderFidProductsOrderProductFidOK), nil

}

/*
GetOrdersOrderFid retrieves an order
*/
func (a *Client) GetOrdersOrderFid(params *GetOrdersOrderFidParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrdersOrderFidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrdersOrderFidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrdersOrderFid",
		Method:             "GET",
		PathPattern:        "/orders/{orderFid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrdersOrderFidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrdersOrderFidOK), nil

}

/*
GetOrdersOrderFidProducts retrieves order products
*/
func (a *Client) GetOrdersOrderFidProducts(params *GetOrdersOrderFidProductsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrdersOrderFidProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrdersOrderFidProductsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrdersOrderFidProducts",
		Method:             "GET",
		PathPattern:        "/orders/{orderFid}/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrdersOrderFidProductsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrdersOrderFidProductsOK), nil

}

/*
PostOrders creates a new order
*/
func (a *Client) PostOrders(params *PostOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*PostOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrdersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOrders",
		Method:             "POST",
		PathPattern:        "/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOrdersOK), nil

}

/*
PostOrdersOrderFidOffers adds an offer to an order
*/
func (a *Client) PostOrdersOrderFidOffers(params *PostOrdersOrderFidOffersParams, authInfo runtime.ClientAuthInfoWriter) (*PostOrdersOrderFidOffersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrdersOrderFidOffersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOrdersOrderFidOffers",
		Method:             "POST",
		PathPattern:        "/orders/{orderFid}/offers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrdersOrderFidOffersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOrdersOrderFidOffersOK), nil

}

/*
PostOrdersOrderFidProducts adds a product to an order
*/
func (a *Client) PostOrdersOrderFidProducts(params *PostOrdersOrderFidProductsParams, authInfo runtime.ClientAuthInfoWriter) (*PostOrdersOrderFidProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrdersOrderFidProductsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOrdersOrderFidProducts",
		Method:             "POST",
		PathPattern:        "/orders/{orderFid}/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrdersOrderFidProductsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOrdersOrderFidProductsOK), nil

}

/*
PutOrdersOrderFidConfirmCard confirms an order authorize the payment
*/
func (a *Client) PutOrdersOrderFidConfirmCard(params *PutOrdersOrderFidConfirmCardParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrdersOrderFidConfirmCardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrdersOrderFidConfirmCardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrdersOrderFidConfirmCard",
		Method:             "PUT",
		PathPattern:        "/orders/{orderFid}/confirmCard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrdersOrderFidConfirmCardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrdersOrderFidConfirmCardOK), nil

}

/*
PutOrdersOrderFidConfirmCoinbase confirms an order await payment
*/
func (a *Client) PutOrdersOrderFidConfirmCoinbase(params *PutOrdersOrderFidConfirmCoinbaseParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrdersOrderFidConfirmCoinbaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrdersOrderFidConfirmCoinbaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrdersOrderFidConfirmCoinbase",
		Method:             "PUT",
		PathPattern:        "/orders/{orderFid}/confirmCoinbase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrdersOrderFidConfirmCoinbaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrdersOrderFidConfirmCoinbaseOK), nil

}

/*
PutOrdersOrderFidConfirmNewCard confirms an order with a new card authorize the payment
*/
func (a *Client) PutOrdersOrderFidConfirmNewCard(params *PutOrdersOrderFidConfirmNewCardParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrdersOrderFidConfirmNewCardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrdersOrderFidConfirmNewCardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrdersOrderFidConfirmNewCard",
		Method:             "PUT",
		PathPattern:        "/orders/{orderFid}/confirmNewCard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrdersOrderFidConfirmNewCardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrdersOrderFidConfirmNewCardOK), nil

}

/*
PutOrdersOrderFidConfirmPayPal confirms an order authorize the payment
*/
func (a *Client) PutOrdersOrderFidConfirmPayPal(params *PutOrdersOrderFidConfirmPayPalParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrdersOrderFidConfirmPayPalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrdersOrderFidConfirmPayPalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrdersOrderFidConfirmPayPal",
		Method:             "PUT",
		PathPattern:        "/orders/{orderFid}/confirmPayPal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrdersOrderFidConfirmPayPalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrdersOrderFidConfirmPayPalOK), nil

}

/*
PutOrdersOrderFidProductsOrderProductFidQuantity sets the quantity of a product on an order
*/
func (a *Client) PutOrdersOrderFidProductsOrderProductFidQuantity(params *PutOrdersOrderFidProductsOrderProductFidQuantityParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrdersOrderFidProductsOrderProductFidQuantityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrdersOrderFidProductsOrderProductFidQuantityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrdersOrderFidProductsOrderProductFidQuantity",
		Method:             "PUT",
		PathPattern:        "/orders/{orderFid}/products/{orderProductFid}/quantity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrdersOrderFidProductsOrderProductFidQuantityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrdersOrderFidProductsOrderProductFidQuantityOK), nil

}

/*
PutOrdersOrderFidVerify verifies an order returning any security urls
*/
func (a *Client) PutOrdersOrderFidVerify(params *PutOrdersOrderFidVerifyParams, authInfo runtime.ClientAuthInfoWriter) (*PutOrdersOrderFidVerifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrdersOrderFidVerifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutOrdersOrderFidVerify",
		Method:             "PUT",
		PathPattern:        "/orders/{orderFid}/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOrdersOrderFidVerifyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOrdersOrderFidVerifyOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}