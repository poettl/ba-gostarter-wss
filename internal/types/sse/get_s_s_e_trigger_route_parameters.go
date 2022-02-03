// Code generated by go-swagger; DO NOT EDIT.

package sse

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetSSETriggerRouteParams creates a new GetSSETriggerRouteParams object
// no default values defined in spec.
func NewGetSSETriggerRouteParams() GetSSETriggerRouteParams {

	return GetSSETriggerRouteParams{}
}

// GetSSETriggerRouteParams contains all the bound params for the get s s e trigger route operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetSSETriggerRoute
type GetSSETriggerRouteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Value
	  Required: true
	  In: path
	*/
	Value string `param:"value"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSSETriggerRouteParams() beforehand.
func (o *GetSSETriggerRouteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rValue, rhkValue, _ := route.Params.GetOK("value")
	if err := o.bindValue(rValue, rhkValue, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSSETriggerRouteParams) Validate(formats strfmt.Registry) error {
	var res []error

	// value
	// Required: true
	// Parameter is provided by construction from the route

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindValue binds and validates parameter Value from path.
func (o *GetSSETriggerRouteParams) bindValue(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Value = raw

	return nil
}
