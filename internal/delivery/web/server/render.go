package server

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

type HTMLRenderer struct {
	layout *template.Template
}

func NewHtmlRenderer() *HTMLRenderer {
	layout := layoutTemplate()
	return &HTMLRenderer{layout}
}

func (t *HTMLRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	fileName := fmt.Sprintf("web/views/pages/%s.html", name)
	temp, err := t.layout.Clone()
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	temp = template.Must(temp.ParseGlob(fileName))

	return temp.ExecuteTemplate(w, name, data)
}

func layoutTemplate() *template.Template {
	funcs := template.FuncMap{
		"StrContains": strings.Contains,
		// "FormatData": time.Now().Format("02 January 2006 - 15:04 WIB"),
	}
	return template.Must(template.New("").Funcs(funcs).ParseGlob("web/views/layouts/*.html"))
}
