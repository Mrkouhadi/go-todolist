package forms

type errors map[string][]string

// Add adds a new error
func (err errors) Add(field, message string) {
	err[field] = append(err[field], message)
}

// Get gets the wanted error
func (err errors) Get(field string) string {
	e := err[field]
	if len(e) == 0 {
		return ""
	}
	return e[0]
}
