package components

import (
	"bytes"
	"text/template"
)

type PostgresPgxPlaceholders struct {
	ProjectPlaceholders
}

func MarshalPostgresPgxTemp(tmpl *template.Template, values interface{}) ([]byte, error) {

	data, ok := values.(PostgresPgxPlaceholders)
	if !ok {
		return nil, nil
	}
	var buf bytes.Buffer
	// exec
	err = tmpl.Execute(&buf, data)

	return buf.Bytes(), nil
}
