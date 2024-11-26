
# sqltmpl 🚀

**`sqltmpl`** is a powerful and flexible SQL template engine for Go that takes the hassle out of writing safe, dynamic SQL statements! Built on top of Go’s `text/template`, `sqltmpl` helps you bind parameters securely and simplifies complex query generation. Perfect for developers who want clean, readable SQL without compromising on safety.

## 🌟 Features

- **Dynamic SQL Binding**: Handles both single values and `IN` clauses with a smart `Bind` function.
- **Flexible Pattern Matching**: Supports `LIKE`, `ILIKE`, and other custom operators with `Concat` for pattern building.
- **Easy-to-Use Templating**: Write your SQL queries in structured, maintainable templates.
- **Cross-Database Support**: Customize placeholder binding to suit PostgreSQL, SQLite, MySQL, and more!

## Installation

```bash
go get github.com/alash3al/sqltmpl
```

## 💡 Quick Start Guide

1. **Define Your SQL Template File (`sql.tmpl`)**:

    ```sql
    {{define "get_user_by_email"}}
        SELECT * FROM users
        WHERE email IN({{.Bind .Args.emails}})
        OR (email = {{.Bind .Args.email}})
        OR email LIKE {{.Bind (.Concat "%" .Args.email "%")}}
    {{end}}
    ```

2. **Use `sqltmpl` in Your Go Code**:

    ```go
    package main

    import (
        "fmt"
        "text/template"
        "github.com/alash3al/sqltmpl"
    )

    func main() {
        // Initialize sqltmpl with template file and PostgreSQL-style bind
        tpl := sqltmpl.New(
            template.Must(template.ParseFiles("sql.tmpl")),
            func(i int) string {
                return fmt.Sprintf("$%d", i) // Bind style for PostgreSQL
            },
        )

        // Execute template with your parameters
        sql, args, err := tpl.Execute("get_user_by_email", map[string]any{
            "emails": []string{"user1@example.com", "user2@example.com"},
            "email":  "user3@example.com",
        })
        if err != nil {
            panic(err)
        }

        fmt.Println(sql, args)
    }
    ```

3. **Output**:

    ```sql
    SELECT * FROM users
    WHERE email IN ($1, $2)
    OR (email = $3)
    OR email LIKE $4
    ```

## 🔧 Advanced Usage

### Smart `Bind` Function for Flexible Binding

`sqltmpl` automatically detects when to use:
- **Single Value Binding**: For cases like `WHERE column = value`.
- **Multi-Value Binding**: For `IN` clauses (`WHERE column IN (value1, value2, ...)`).

### `Concat` Function for Pattern Matching

Build dynamic patterns with `%` wildcards for `LIKE` and `ILIKE` searches:

```sql
WHERE column LIKE {{.Bind (.Concat "%" .Args.pattern "%")}}
```

### Custom Binding Styles for Your Database

`sqltmpl` lets you define custom binding styles for any SQL engine:

```go
tpl := sqltmpl.New(myTemplate, func(i int) string { return fmt.Sprintf(":%d", i) }) // e.g., Oracle-style
```

## 📚 Example Use Cases

- **User Search with `LIKE` or `ILIKE`**: Perform case-insensitive searches in SQL databases.
- **Multi-Value Filters**: Filter records with multiple values using `IN` clauses.
- **Dynamic SQL Generation**: Build complex SQL statements using conditions, ranges, and subqueries.

## 💡 Why Choose `sqltmpl`?

- **Safety First**: Prevents SQL injection with automatic binding for both single values and slices.
- **Clean and Reusable SQL**: Manage all your SQL in template files, keeping code organized and readable.
- **Database Agnostic**: Works seamlessly across different databases with customizable placeholders.

## 🛠️ Contributing

We welcome contributions! Whether it’s a new feature, bug fix, or documentation improvement, feel free to open a PR or issue. Let's make SQL templating easier for everyone! ✨

## 📄 License

`sqltmpl` is licensed under the MIT License.
