package files

import (
	"io"
	"strings"
	"text/template"

	"github.com/fatih/camelcase"
)

func RunTemplate(f io.Writer, templateStr string, obj interface{}) error {
	funcs := template.FuncMap{
		"Title": func(s string) string {
			return strings.Title(s)
		},
		"camelcaseSplit": func(s string) string { return strings.Join(camelcase.Split(s), " ") },
	}

	tmpl, err := template.New("").Funcs(funcs).Parse(templateStr)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, obj); err != nil {
		return err
	}

	return nil
}
