package render

import (
	"bytes"
	"html/template"
	"intelchaos/pkg/config"
	"intelchaos/pkg/models"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig
func NewTemplates(a *config.AppConfig){
	app = a
}

//RenderTemplate renders templates using html
var tc map[string]*template.Template

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData){
	//create a template cache
	if app.UseCache{
		tc = app.TemplateCache
	}else{
		tc,_ = CreateTemplateCache()
	}


	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok{
		log.Fatal("Could not get Template")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	//render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("error parsing template:", err)
	}
}

func CreateTemplateCache() (map[string] *template.Template, error)	{
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl from templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil{
		 return myCache, err
	}

	//range through all of the files
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil{
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil{
			return myCache, err
		}
		if len(matches) > 0{
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}
		myCache[name] = ts
	}
	return myCache, nil
}
