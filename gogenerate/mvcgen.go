package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

const version = "0.0.1 00001"

var workdir = "."

var DIV string = "\n//DO NOT CHANGE ABOVE --GENERATED--"

func main() {
	fmt.Printf("version %s\n", version)
	fmt.Printf("workdir %s\n", workdir)
	for i, name := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf("%d gen resource %s\n", i, name)
		genFile(workdir+string(os.PathSeparator)+"controllers", name, "controller", controllerTemplate, makeTemplateFile)
		genFile(workdir+string(os.PathSeparator)+"views", name, "view", viewTemplate, makeTemplateFile)
	}
}
func genFile(dir string, name string, suff string, template string, maker func(name string, filecontents string, texttemplate string, w io.Writer) (err error)) {
	var path = fmt.Sprintf("%s%s%s_%s", dir, string(os.PathSeparator), name, suff)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0775)
		fmt.Printf("making dir %s\n", path)
	}
	path = fmt.Sprint(dir, string(os.PathSeparator), name, "_", suff, string(os.PathSeparator), name, "_", suff, ".go")
	var filecontents string
	if _, err := os.Stat(path); os.IsNotExist(err) {
		filecontents = ""
		fmt.Printf("file created %s\n", path)
	} else {
		b, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		filecontents = string(b)
	}
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file modified %s\n", path)
	defer f.Close()
	err = maker(name, filecontents, template, f)
	if err != nil {
		panic(err)
	}

}

var controllerTemplate string = `package {{.Name}}_controller
//import{{ if .Import -}} {{ .Import}} {{ else }} 
import (
	"github.com/fops9311/mvc_server_app/model/controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/views/{{.Name}}_view"
){{ end }}//import
var Resource resource.Resurce
var new_{{.Name}} controller.Action = func(params map[string]string) (result string, err error) {return "", nil}
var get_{{.Name}}s controller.Action = func(params map[string]string) (result string, err error) {return "", nil}
var get_{{.Name}}_by_id controller.Action = func(params map[string]string) (result string, err error) {return "", nil}
var update_{{.Name}}_by_id controller.Action = func(params map[string]string) (result string, err error) {return "", nil}
var delete_{{.Name}}_by_id controller.Action = func(params map[string]string) (result string, err error) {return "", nil}

func init() {
	init_begin()
	Resource = resource.NewResource()
	Resource.Key = "/{{.Name}}"
	Resource.Actions = []resource.ActionPath{
		{
			Verb:       "POST",
			Path:       "/",
			Middleware: make([]string, 0),
			Action:     new_{{.Name}},
		},
		{
			Verb:       "GET",
			Path:       "/",
			Middleware: make([]string, 0),
			Action:     get_{{.Name}}s,
		},
		{
			Verb:       "GET",
			Path:       "/:{{.Name}}_id",
			Middleware: make([]string, 0),
			Action:     get_{{.Name}}_by_id,
		},
		{
			Verb:       "PUT",
			Path:       "/:{{.Name}}_id",
			Middleware: make([]string, 0),
			Action:     update_{{.Name}}_by_id,
		},
		{
			Verb:       "DELETE",
			Path:       "/:{{.Name}}_id",
			Middleware: make([]string, 0),
			Action:     delete_{{.Name}}_by_id,
		},
	}
	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}`

var viewTemplate string = `package {{.Name}}_view
//import{{ if .Import -}} {{ .Import}} {{ else }} 
import (
	"fmt"
){{ end }}//import
func Dummy(){
	fmt.Print("hi")
}`

func makeTemplateFile(name string, filecontents string, texttemplate string, w io.Writer) (err error) {
	data := struct {
		Name   string
		Import string
	}{
		Name:   name,
		Import: "",
	}
	if importpart := getImportPart(filecontents); len(importpart) > 0 {
		data.Import = importpart
	}
	sects := replaceGenPart(texttemplate, filecontents)
	tmpl, err := template.New(name).Parse(sects[0])
	defer func() {
		for i, sect := range sects {
			if i > 0 {
				w.Write([]byte(sect))
			}
		}
	}()
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

func replaceGenPart(genpart string, filecontents string) (result []string) {
	s := strings.Split(filecontents, DIV)
	if len(s) < 2 {
		return append(s, DIV)
	}
	s[0] = genpart + DIV
	return s
}
func getImportPart(filecontents string) (result string) {
	s := strings.Split(filecontents, "//import")
	if len(s) > 2 {
		return s[1]
	}
	return ""
}
