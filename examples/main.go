package main

import (
	"fmt"
	"github.com/alash3al/sqltmpl"
	"text/template"
)

func main() {
	tpl := sqltmpl.New(
		template.Must(template.ParseFiles("sql.tmpl")),
		func(i int) string {
			return fmt.Sprintf("$%d", i)
		},
	)

	fmt.Println(tpl.Execute("get_user_by_email", map[string]any{
		"email":  "m@e.com",
		"emails": []string{"e1@o.com", "e2@o.com"},
	}))
}
