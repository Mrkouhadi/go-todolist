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
	if xfield == "" {
		f.Errors.Add(field, "This field cannot be Blank !")
		return false
	}
	return true
}

// Valid returns true if there are no errors, false if there are errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
