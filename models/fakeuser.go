package models

import (
	"os"
	"text/template"
)

type FakeUser struct {
	Name     string
	Address  string
	SSN      string
	Phone    string
	Birthday string
	Email    string
	Username string
	Password string
	Height   string
	Weight   string
}

var tpl = `Name: {{ .Name }}
Address: {{ .Address }}

SSN: {{ .SSN }}

Phone: {{ .Phone }}
Birthday: {{ .Birthday }}
Email: {{ .Email }}
Username: {{ .Username }}
Password: {{ .Password }}

Height: {{ .Height }}
Weight: {{ .Weight }}
`

func (f *FakeUser) Save() {
	t, _ := template.New("").Parse(tpl)
	t.Execute(os.Stdout, f)

	file, _ := os.Create(f.Username)
	defer file.Close()

	t.Execute(file, f)
}
