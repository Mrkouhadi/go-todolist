package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

// New initalizes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if a field is empty
func (f *Form) Has(field string, r *http.Request) bool {
	xfield := r.Form.Get(field)
	return xfield != ""
}
