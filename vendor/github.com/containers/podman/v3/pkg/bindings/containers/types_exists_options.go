// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/containers/podman/v3/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *ExistsOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *ExistsOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithExternal set field External to given value
func (o *ExistsOptions) WithExternal(value bool) *ExistsOptions {
	o.External = &value
	return o
}

// GetExternal returns value of field External
func (o *ExistsOptions) GetExternal() bool {
	if o.External == nil {
		var z bool
		return z
	}
	return *o.External
}
