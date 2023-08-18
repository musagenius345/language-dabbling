package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Email   string
	IsAdmin bool
}

func main() {
	tmpl := template.Must(template.New("user").Parse(`
		<h1>Hello, {{.Name}}!</h1>
		<p>Email: {{.Email}}</p>
		{{if .IsAdmin}}
			<p>You are an admin.</p>
		{{else}}
			<p>You are not an admin.</p>
		{{end}}
	`))

	user := User{Name: "Alice", Email: "alice@example.com", IsAdmin: true}
	err := tmpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
