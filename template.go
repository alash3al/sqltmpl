package sqltmpl

import (
	"bytes"
	"strings"
	"text/template"
)

// Template is the representation of a parsed template.
type Template struct {
	tpl             *template.Template
	placeholderFunc func(i int) string
}

func New(tpl *template.Template, placeholderFunc func(i int) string) *Template {
	return &Template{
		tpl:             tpl,
		placeholderFunc: placeholderFunc,
	}
}

// Execute executes a template by the specified name and pass the specified data to the template Context
// then it returns the (SQL statement, the bindings) or any error if any.
func (t *Template) Execute(name string, data any) (string, []any, error) {
	output := bytes.NewBuffer(nil)
	ctx := Context{Args: data, placeholderFunc: t.placeholderFunc}

	if err := t.tpl.ExecuteTemplate(output, name, &ctx); err != nil {
		return "", nil, err
	}

	return strings.TrimSpace(output.String()), ctx.bindings, nil
}
