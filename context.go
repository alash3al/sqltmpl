package sqltmpl

import (
	"reflect"
	"strings"
)

// Context is what is being exposed to the template being executed
// it contains the args passed while executing a template.
type Context struct {
	Args            any
	bindings        []any
	placeholderFunc func(i int) string
}

// Bind a helper function that safely injects a value into the SQL statement
//
// Example `SELECT * FROM users WHERE email = {{.Bind .Args.email}}`
func (c *Context) Bind(value any) string {
	v := reflect.ValueOf(value)

	if (v.Kind() == reflect.Slice) || (v.Kind() == reflect.Array) {
		var result []string

		for i := 0; i < v.Len(); i++ {
			result = append(result, c.Bind(v.Index(i).Interface()))
		}

		return strings.Join(result, ", ")
	}

	c.bindings = append(c.bindings, value)
	return c.placeholderFunc(len(c.bindings))
}

// Concat a helper function that safely injects multiple values into the SQL statement
//
// Example `SELECT * FROM users WHERE email in ({{.Concat .Args.emails}})`
func (c *Context) Concat(values ...string) string {
	return strings.Join(values, "")
}
