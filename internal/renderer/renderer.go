package renderer

import (
	"bytes"
	"embed"
	"strings"
	"text/template"

	m "github.com/0xspiderr/gododocs/internal/model"
)

//go:embed templates/*.tmpl
var tmplFS embed.FS

// purpose: Joins parameters by name and type and then separates
// them by comma. This method makes templates lighter. It is
// necessary also because both Method[] and Signal[] have []Param
// and I didnt want to repeat the code for both of them.
func joinParameters(params []m.Param) string {
	pr := make([]string, len(params))
	for i, p := range params {
		pr[i] += p.Name + ": " + p.Type
	}
	return strings.Join(pr, ", ")
}

// purpose: Read the embedded templates dir, concat
// them with the file names in that dir and return the
// string slice. This is a helper used in the Render() func
// when building a new template.
func classFiles() []string {
	files, err := tmplFS.ReadDir("templates")
	if err != nil {
		panic(err)
	}
	names := make([]string, 0, len(files))
	for _, f := range files {
		names = append(names, "templates/"+f.Name())
	}

	return names
}

// purpose: Apply the templates on a Class structure and
// return the resulting bytes to be written to a file
// by writer.go
func Render(c m.Class) (string, error) {
	fm := template.FuncMap{"joinParameters": joinParameters}
	t, err := template.New("main.tmpl").Funcs(fm).ParseFS(tmplFS, classFiles()...)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, c); err != nil {
		return "", err
	}
	return buf.String(), nil
}
