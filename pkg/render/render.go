package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rhs99/reserveit/pkg/config"
	"github.com/rhs99/reserveit/pkg/models"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData)*models.TemplateData{
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData){
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	}else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template!")
	}

	buff := new(bytes.Buffer)
	
	td = AddDefaultData(td)

	_ = t.Execute(buff, td)

	_, err := buff.WriteTo(w)

	if err != nil{
		fmt.Println("Error writing template", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return tmplCache, err
	}

	for _, page := range pages{
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return tmplCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err!=nil{
			return tmplCache, err;
		}

		if len(matches) > 0{
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return tmplCache, err
			}
		}

		tmplCache[name] = ts
	}
	return tmplCache, nil
}