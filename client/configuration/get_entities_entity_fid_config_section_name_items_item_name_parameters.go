// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParams creates a new GetEntitiesEntityFidConfigSectionNameItemsItemNameParams object
// with the default values initialized.
func NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParams() *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	var ()
	return &GetEntitiesEntityFidConfigSectionNameItemsItemNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParamsWithTimeout creates a new GetEntitiesEntityFidConfigSectionNameItemsItemNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParamsWithTimeout(timeout time.Duration) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	var ()
	return &GetEntitiesEntityFidConfigSectionNameItemsItemNameParams{

		timeout: timeout,
	}
}

// NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParamsWithContext creates a new GetEntitiesEntityFidConfigSectionNameItemsItemNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParamsWithContext(ctx context.Context) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	var ()
	return &GetEntitiesEntityFidConfigSectionNameItemsItemNameParams{

		Context: ctx,
	}
}

// NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParamsWithHTTPClient creates a new GetEntitiesEntityFidConfigSectionNameItemsItemNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEntitiesEntityFidConfigSectionNameItemsItemNameParamsWithHTTPClient(client *http.Client) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	var ()
	return &GetEntitiesEntityFidConfigSectionNameItemsItemNameParams{
		HTTPClient: client,
	}
}

/*GetEntitiesEntityFidConfigSectionNameItemsItemNameParams contains all the parameters to send to the API endpoint
for the get entities entity fid config section name items item name operation typically these are written to a http.Request
*/
type GetEntitiesEntityFidConfigSectionNameItemsItemNameParams struct {

	/*EntityFid
	  Entity FID to use

	*/
	EntityFid string
	/*ItemName
	  Item Name

	*/
	ItemName string
	/*SectionName
	  Section Name

	*/
	SectionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) WithTimeout(timeout time.Duration) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) WithContext(ctx context.Context) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) WithHTTPClient(client *http.Client) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityFid adds the entityFid to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) WithEntityFid(entityFid string) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	o.SetEntityFid(entityFid)
	return o
}

// SetEntityFid adds the entityFid to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) SetEntityFid(entityFid string) {
	o.EntityFid = entityFid
}

// WithItemName adds the itemName to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) WithItemName(itemName string) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	o.SetItemName(itemName)
	return o
}

// SetItemName adds the itemName to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) SetItemName(itemName string) {
	o.ItemName = itemName
}

// WithSectionName adds the sectionName to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) WithSectionName(sectionName string) *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams {
	o.SetSectionName(sectionName)
	return o
}

// SetSectionName adds the sectionName to the get entities entity fid config section name items item name params
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) SetSectionName(sectionName string) {
	o.SectionName = sectionName
}

// WriteToRequest writes these params to a swagger request
func (o *GetEntitiesEntityFidConfigSectionNameItemsItemNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityFid
	if err := r.SetPathParam("entityFid", o.EntityFid); err != nil {
		return err
	}

	// path param itemName
	if err := r.SetPathParam("itemName", o.ItemName); err != nil {
		return err
	}

	// path param sectionName
	if err := r.SetPathParam("sectionName", o.SectionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
