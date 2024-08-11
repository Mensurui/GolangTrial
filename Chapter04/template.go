package main

import (
	"html/template"
	"log"
	"os"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	// Example data for testing
	data := struct {
		TotalCount int
		Items      []struct {
			HTMLURL string
			Number  int
			State   string
			User    struct {
				HTMLURL string
				Login   string
			}
			Title string
		}
	}{
		TotalCount: 3,
		Items: []struct {
			HTMLURL string
			Number  int
			State   string
			User    struct {
				HTMLURL string
				Login   string
			}
			Title string
		}{
			{"https://example.com/issue/1", 1, "open", struct {
				HTMLURL string
				Login   string
			}{"https://example.com/user/1", "user1"}, "Issue 1"},
			{"https://example.com/issue/2", 2, "closed", struct {
				HTMLURL string
				Login   string
			}{"https://example.com/user/2", "user2"}, "Issue 2"},
			{"https://example.com/issue/3", 3, "open", struct {
				HTMLURL string
				Login   string
			}{"https://example.com/user/3", "user3"}, "Issue 3"},
		},
	}

	// Render the template with data
	if err := issueList.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
