package render

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

// renders and writes templats based on given tmpl name
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//root project path
	WORKPATH := "/work/"

	//caching the templates in templateCache map
	templateCache, err := formTemplateCache(WORKPATH)

	if err != nil {
		fmt.Println("error: template cache is not available, loading from disk")
		//loading the specific file into cache as the file is not found in the cache
		templateCache[tmpl], err = template.ParseFiles(WORKPATH+"templates/"+tmpl, WORKPATH+"templates/base.layout.tmpl")

		if err != nil {
			fmt.Println("error: loading", tmpl, "beacuse", err)
			return
		}
	}

	//loading the file from cache
	parsedTemplate, fileExists := templateCache[tmpl]
	if !fileExists {
		fmt.Println("error: template file doesnot exsit", tmpl)
		return
	}

	//writing the file in http response writer
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error rendering template: ", tmpl, "with error:", err)
		return
	}
}

func formTemplateCache(WORKPATH string) (map[string]*template.Template, error) {
	fmt.Println("verbose: loading template cache")

	tc := map[string]*template.Template{}

	files, err := os.ReadDir(WORKPATH + "templates/")
	if err != nil {
		fmt.Println("error: cannot read files at WORKPATH,", err)
		return nil, err
	}

	for _, file := range files {
		tc[file.Name()], err = template.ParseFiles(WORKPATH+"templates/"+file.Name(), WORKPATH+"templates/base.layout.tmpl")
		if err != nil {
			fmt.Println("error: cannot read file", WORKPATH+"templates/"+file.Name())
			return nil, err
		}
	}
	return tc, nil
}
