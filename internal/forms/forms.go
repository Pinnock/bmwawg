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
	return r.Form.Get(field) != ""
}
