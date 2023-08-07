package render

import (
	"text/template"
)

func Rende(templateContent string, args map[string]interface{}) (string, error) {
	// use template and args to render
	name := "test"
	template.New(name)
	return "", nil
}
