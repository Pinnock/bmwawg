package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
		return false
	}
	return true
}

func (i *Info) Require(fields ...string) {
	for _, f := range fields {
		if strings.TrimSpace(i.Get(f)) == "" {
			i.Errors.Add(f, "This is a required field")
		}
	}
}

func (i *Info) HasMinLength(field string, length int, r *http.Request) bool {
	if len(r.Form.Get(field)) < length {
		i.Errors.Add(
			field,
			fmt.Sprintf("This field must be at least %d characters long", length),
		)
		return false
	}

	return true
}
func (f *Info) Valid() bool {
	return len(f.Errors) == 0
}
