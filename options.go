package api

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WithDisableNesting disables automatic expansion of related FIDs for one call.
// Pass it as an optional argument to any generated client method. The server
// must support the Disable-Nesting header.
func WithDisableNesting(operation *runtime.ClientOperation) {
	withNestingHeader(operation, "Disable-Nesting")
}

// WithShortNesting requests compact related objects (id, fid, and displayName)
// for one call. Pass it as an optional argument to any generated client method.
// The server must support the Short-Nesting header. Disable-Nesting takes
// precedence on the backend if both options are supplied.
func WithShortNesting(operation *runtime.ClientOperation) {
	withNestingHeader(operation, "Short-Nesting")
}

func withNestingHeader(operation *runtime.ClientOperation, header string) {
	params := operation.Params
	operation.Params = runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, formats strfmt.Registry) error {
		if params != nil {
			if err := params.WriteToRequest(req, formats); err != nil {
				return err
			}
		}
		return req.SetHeaderParam(header, "true")
	})
}
