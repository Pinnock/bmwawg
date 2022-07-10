package forms

type errors map[string][]string

func (e errors) Add(field, errmsg string) {
	e[field] = append(e[field], errmsg)
}

func (e errors) Get(field string) []string {
	_, ok := e[field]
	if !ok {
		return []string{}
	}
	return e[field]
}
