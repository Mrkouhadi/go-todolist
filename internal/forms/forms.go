package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
		return false
	}
	return true
}

// Valid returns true if there are no errors, false if there are errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required checks reqired fields and add errors when they're empty
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		val := f.Get(field)
		if strings.TrimSpace(val) == "" {
			f.Errors.Add(field, "This Field cannot be Blank !")
		}
	}
}

// MinLength checks fr the fields minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	xfield := r.Form.Get(field)
	if len(xfield) < length {
		f.Errors.Add(field, fmt.Sprintf("This field should be at least characters %d long", length))
		return false
	}
	return true
}
