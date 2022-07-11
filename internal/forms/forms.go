package forms

import (
	"net/http"
	"net/url"
)

type Info struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Info {
	return &Info{
		Values: data,
		Errors: errors{},
	}
}

func (i *Info) HasValue(field string, r *http.Request) bool {
	if r.Form.Get(field) == "" {
		i.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

func (f *Info) Valid() bool {
	return len(f.Errors) == 0
}
