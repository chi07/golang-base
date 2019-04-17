package helpers


import (
	"html/template"
	"strings"
)

var templates map[string]*template.Template

func init()  {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
}

/**
 Return Substr with correct format
 */
func SubStr(str string, len int) string {
	trimStr := strings.TrimSpace(str)

	//runes := []rune(trimStr)
	//return string(runes[:len])
	return string(trimStr)
}